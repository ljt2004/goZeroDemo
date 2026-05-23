package logic

import (
	"context"
	"strconv"
	"time"

	"user-grpc/internal/model"
	"user-grpc/internal/svc"
	"user-grpc/internal/utils"
	user_grpc "user-grpc/user-grpc"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user_grpc.LoginRequest) (*user_grpc.LoginResponse, error) {

	DB := l.svcCtx.DB

	var u model.User
	if DB.Where("phone = ?", in.Phone).First(&u).Error != nil {
		return nil, status.Error(400, "手机号不存在")
	}

	if err := utils.CheckPassword(in.Password, u.Password); err == false {
		logx.Errorf("密码错误")
		return nil, status.Error(400, "密码错误")
	}

	// 3. 生成 JWT（用 go-zero 官方同款方式）
	secret := l.svcCtx.Config.JwtAuth.AccessSecret // 从配置读
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	now := time.Now().Unix()

	claims := jwt.MapClaims{
		"userId": strconv.FormatInt(u.ID, 10), //转为字符串,否则会失精度
		"iat":    now,
		"exp":    now + expire,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return nil, status.Error(codes.Internal, "token生成失败")
	}

	return &user_grpc.LoginResponse{
		Token: token,
	}, nil
}
