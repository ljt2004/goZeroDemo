package logic

import (
	"context"

	"user-grpc/internal/svc"
	"user-grpc/user-grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user_grpc.GetUserRequest) (*user_grpc.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user_grpc.GetUserResponse{}, nil
}
