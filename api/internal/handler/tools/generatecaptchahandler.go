// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tools

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroApi/internal/logic/tools"
	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"
)

// 获取图形验证码
func GenerateCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tools.NewGenerateCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GenerateCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
