package logic

import (
	"context"

	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackOrderLogic {
	return &RollbackOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 回滚订单
func (l *RollbackOrderLogic) RollbackOrder(in *order.CreateOrderReq) (*order.CreateOrderRsp, error) {
	// todo: add your logic here and delete this line

	return &order.CreateOrderRsp{}, nil
}
