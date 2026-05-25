// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tools

import (
	"context"
	"strings"

	"goZeroApi/internal/pkg"
	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取图形验证码
func NewGenerateCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCaptchaLogic {
	return &GenerateCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取图形验证码
// 采用 github.com/mojocn/base64Captcha  生成验证码

func (l *GenerateCaptchaLogic) GenerateCaptcha(req *types.GenerateCaptchaReq) (resp *types.GenerateCaptchaResp, err error) {

	// 1. 创建数字验证码驱动
	driver := base64Captcha.NewDriverDigit(
		l.svcCtx.Config.Captcha.Height,
		l.svcCtx.Config.Captcha.Width,
		l.svcCtx.Config.Captcha.Length,
		l.svcCtx.Config.Captcha.MaxSkew,
		20)

	// 2. 使用 Redis 存储
	store := pkg.NewRedisStore(l.svcCtx.Redis)

	// 3. 创建验证码实例
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 4. 生成验证码
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		l.Error("生成验证码失败:", err)
		return nil, err
	}

	// 5. 处理 base64 字符串
	base64Code := strings.Replace(b64s, "data:image/png;base64,", "", 1)

	// 6. 返回给前端
	resp = &types.GenerateCaptchaResp{
		Id:    id,
		Image: base64Code,
	}

	return resp, nil
}
