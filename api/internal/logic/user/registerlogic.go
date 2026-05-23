// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"
	user_grpc "user-grpc/user-grpc"

	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResponse, err error) {
	// 统一返回结构体，不管成功失败都用这个
	resp = &types.BaseResponse{
		Code:    0,
		Message: "success",
		Data:    nil,
	}

	// 1. 参数校验（错误直接设置 resp，不返回 err）
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
	if req.Username == "" {
		l.Logger.Errorf("username is empty")
		resp.Code = 400
		resp.Message = "用户名不能为空"
		return resp, nil
	}

	// 2. 调用RPC
	grpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user_grpc.RegisterRequest{
		Phone:    req.Phone,
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		// 直接从 gRPC 错误里提取消息
		st, ok := status.FromError(err)
		if ok {
			resp.Message = st.Message() // 直接得到：手机号已存在
		} else {
			resp.Message = err.Error()
		}
		resp.Code = 400
		return resp, nil
	}

	// 3. 成功
	resp.Code = 200
	resp.Message = "注册成功"
	resp.Data = grpcResp

	return resp, nil
}
