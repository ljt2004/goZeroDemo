// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tools

import (
	"net/http"
	"strings"

	"goZeroApi/internal/logic/tools"
	"goZeroApi/internal/pkg/response"
	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendEmailCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendEmailCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Fail(w, http.StatusBadRequest, "请求参数错误")
			return
		}

		if strings.TrimSpace(req.Email) == "" {
			response.Fail(w, http.StatusBadRequest, "邮箱不能为空")
			return
		}

		l := tools.NewSendEmailCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendEmailCode(&req)

		if err != nil {
			response.Fail(w, http.StatusInternalServerError, err.Error())
		} else {
			response.Success(w, resp)
		}
	}
}
