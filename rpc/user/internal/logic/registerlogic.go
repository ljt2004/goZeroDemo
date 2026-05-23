package logic

import (
	"context"

	"user-grpc/internal/model"
	"user-grpc/internal/svc"
	"user-grpc/internal/utils"
	user_grpc "user-grpc/user-grpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user_grpc.RegisterRequest) (*user_grpc.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	// 注册逻辑
	if in.Username == "" || in.Password == "" {
		logx.Errorf("用户名或密码不能为空")
		return nil, status.Error(codes.InvalidArgument, "用户名或密码不能为空")
	}

	if in.Phone == "" {
		logx.Errorf("手机号不能为空")
		return nil, status.Error(codes.InvalidArgument, "手机号不能为空")
	}

	if len(in.Phone) != 11 {
		logx.Errorf("手机号长度不正确")
		return nil, status.Error(codes.InvalidArgument, "手机号长度不正确")
	}
	var DB = l.svcCtx.DB

	var count int64
	DB.Model(&model.User{}).Where("phone = ?", in.Phone).Count(&count)
	if count > 0 {
		logx.Errorf("手机号已存在")
		return nil, status.Error(codes.AlreadyExists, "手机号已存在")
	}

	hashPassword, err := utils.HashPassword(in.Password)
	if err != nil {
		logx.Errorf("密码加密失败")
		return nil, status.Error(codes.Internal, "密码加密失败")
	}
	user := model.User{
		BaseModel: model.BaseModel{
			ID: l.svcCtx.Snowflake.Generate().Int64(),
		},
		Username: in.Username,
		Password: hashPassword,
		Phone:    in.Phone,
	}

	if l.svcCtx.UserDao.CreateUser(&user) != nil {
		logx.Errorf("注册失败")
		return nil, status.Error(codes.Internal, "注册失败")
	}

	logx.Infof("用户名：%s, 密码：%s", in.Username, in.Password)

	return &user_grpc.RegisterResponse{
		Id:       user.ID,
		Username: in.Username,
		Status:   "注册成功",
	}, nil
}
