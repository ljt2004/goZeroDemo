// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tools

import (
	"context"

	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"
	"goZeroApi/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(req *types.SendEmailCodeReq) (resp *types.BaseResponse, err error) {
	resp = &types.BaseResponse{}

	// 1. 创建邮箱发送器
	emailConfig := l.svcCtx.Config.Email
	sender := utils.NewEmailSender(
		emailConfig.Username,
		"GoZeroDemo",     // 发件人昵称
		emailConfig.Host, // 邮件服务器地址
		emailConfig.Port, // SMTP端口号
		emailConfig.Password,
		true, // 使用 TLS
	)

	// 2. 创建验证码服务
	service := utils.NewEmailCodeService(
		utils.NewNumberCodeGenerator(6), // 6位数字验证码
		utils.NewRedisCodeStore(l.svcCtx.Redis, "email_code:"),
		sender,
		300, // 5分钟过期
		`<h3>你的验证码是：%s</h3><p>%d分钟内有效，请勿泄露</p>`,
	)

	// 3. 发送验证码
	err = service.SendCode(req.Email)
	if err != nil {
		l.Error("发送邮箱验证码失败:", err)
		resp.Code = 500
		resp.Message = "发送验证码失败"
		return resp, nil
	}

	resp.Code = 200
	resp.Message = "验证码发送成功"
	return resp, nil
}
