package logic

import (
	"context"

	"user-grpc/internal/model"
	"user-grpc/internal/svc"
	user_grpc "user-grpc/user-grpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetMyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDetailLogic {
	return &GetMyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyDetailLogic) GetMyDetail(in *user_grpc.GetMyDetailRequest) (*user_grpc.GetMyDetailResponse, error) {
	logx.Info("GetMyDetail in: ", in.UserId)
	var user *model.User

	if user, _ = l.svcCtx.UserDao.GetUserById(in.UserId); user == nil {
		logx.Error("用户不存在")
		return nil, status.Errorf(400, "用户不存在")
	}

	return &user_grpc.GetMyDetailResponse{
		Name:  user.Username,
		Phone: user.Phone,
		Email: user.Email,
	}, nil

}
