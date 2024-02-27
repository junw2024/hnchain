package logic

import (
	"context"

	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderCheckLogic {
	return &CreateOrderCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单创建验证
func (l *CreateOrderCheckLogic) CreateOrderCheck(in *order.CreateOrderReq) (*order.CreateOrderRsp, error) {
	// todo: add your logic here and delete this line

	return &order.CreateOrderRsp{}, nil
}
