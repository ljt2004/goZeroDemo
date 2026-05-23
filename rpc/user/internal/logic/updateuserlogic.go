package logic

import (
	"context"

	"user-grpc/internal/svc"
	"user-grpc/user-grpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user_grpc.UpdateUserRequest) (*user_grpc.UpdateUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user_grpc.UpdateUserResponse{}, nil
}
