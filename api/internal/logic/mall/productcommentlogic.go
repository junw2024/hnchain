package mall

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentLogic {
	return &ProductCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentLogic) ProductComment(req *types.ProductCommentReq) (resp *types.ProductCommentRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
