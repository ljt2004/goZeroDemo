// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroApi/internal/logic/article"
	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"
)

func ArticleDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewArticleDetailLogic(r.Context(), svcCtx)
		resp, err := l.ArticleDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
