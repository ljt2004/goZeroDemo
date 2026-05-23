// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"context"

	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDetailLogic {
	return &ArticleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleDetailLogic) ArticleDetail(req *types.ArticleReq) (resp *types.ArticlerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
