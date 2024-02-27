package order

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderHanderLogic {
	return &AddOrderHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrderHanderLogic) AddOrderHander(req *types.OrderAddReq) (resp *types.OrderAddRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
