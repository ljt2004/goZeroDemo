// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"
	user_grpc "user-grpc/user-grpc"

	"goZeroApi/internal/pkg"
	"goZeroApi/internal/svc"
	"goZeroApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetMyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDetailLogic {
	return &GetMyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyDetailLogic) GetMyDetail(req *types.GetMyDetailReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.BaseResponse{
		Code:    200,
		Message: "ok",
		Data:    nil,
	}

	// 这里不需要做错误判断,前jwt中间件已经校验过了，直接获取即可
	// uidStr := l.ctx.Value("userId")

	// logx.Info("uidStr:  ", uidStr)

	// logx.Info("uidStr:  ", uidStr)
	// logx.Info("11111111userId:  ", userId)
	userId, _ := pkg.GetUserIdFromCtx(l.ctx)

	data, err := l.svcCtx.UserRpc.GetMyDetail(l.ctx, &user_grpc.GetMyDetailRequest{
		UserId: userId,
	})

	if err != nil {
		// 直接从 gRPC 错误里提取消息
		st, ok := status.FromError(err)
		if ok {
			resp.Message = st.Message()
		} else {
			resp.Message = err.Error()
		}
		resp.Code = 400
		return resp, nil
	}

	resp.Data = map[string]interface{}{
		"id":    userId,
		"user":  data.Name,
		"phone": data.Phone,
		"email": data.Email,
	}

	return resp, nil
}
