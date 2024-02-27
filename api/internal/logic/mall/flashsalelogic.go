package mall

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashSaleLogic {
	return &FlashSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashSaleLogic) FlashSale() (resp *types.FlashSaleRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
