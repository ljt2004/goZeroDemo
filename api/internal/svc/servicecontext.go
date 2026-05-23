// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"goZeroApi/internal/config"
	user_grpc "user-grpc/user-grpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user_grpc.UserClient

	Authority string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user_grpc.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}
