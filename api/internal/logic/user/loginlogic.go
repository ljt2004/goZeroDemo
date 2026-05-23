// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"
	user_grpc "user-grpc/user-grpc"

	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	resp = &types.BaseResponse{
		Code:    200,
		Message: "success",
		Data:    nil,
	}
	if req.Phone == "" || len(req.Phone) != 11 {
		l.Logger.Errorf("phone is error, phone: %s", req.Phone)
		resp.Code = 400
		resp.Message = "手机号格式不正确"
		return resp, nil // 这里不返回 err！
	}

	if req.Password == "" {
		l.Logger.Errorf("password is empty")
		resp.Code = 400
		resp.Message = "密码不能为空"
		return resp, nil
	}

	LoginResponse, err := l.svcCtx.UserRpc.Login(l.ctx, &user_grpc.LoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})

	if err != nil {
		resp.Code = 400
		resp.Message = "登录失败"
		return resp, nil
	}

	resp.Data = map[string]interface{}{
		"token": LoginResponse.Token,
	}
	resp.Message = "登录成功"

	// Authorization : 前端需要传
	return resp, nil
}
