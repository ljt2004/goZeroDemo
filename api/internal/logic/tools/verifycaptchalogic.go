// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tools

import (
	"context"

	"goZeroApi/internal/pkg"
	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 校验图形验证码
func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCaptchaLogic) VerifyCaptcha(req *types.VerifyCaptchaReq) (resp *types.BaseResponse, err error) {
	resp = &types.BaseResponse{}
	resp.Code = 200
	resp.Message = "验证码正确"
	resp.Data = nil

	// 检查 Redis 是否可用
	if l.svcCtx.Redis == nil {
		resp.Code = 500
		resp.Message = "验证码服务未初始化"
		return resp, nil
	}

	store := pkg.NewRedisStore(l.svcCtx.Redis)

	if !store.Verify(req.Id, req.Code, true) {
		resp.Code = 400
		resp.Message = "验证码错误"
		return resp, nil
	}

	return resp, nil
}
