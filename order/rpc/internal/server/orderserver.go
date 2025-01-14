// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"hnchain/order/rpc/internal/logic"
	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

// 个人订单分页
func (s *OrderServer) Orders(ctx context.Context, in *order.OrdersReq) (*order.OrdersRsp, error) {
	l := logic.NewOrdersLogic(ctx, s.svcCtx)
	return l.Orders(in)
}

// 创建订单
func (s *OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderRsp, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

// 订单创建验证
func (s *OrderServer) CreateOrderCheck(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderRsp, error) {
	l := logic.NewCreateOrderCheckLogic(ctx, s.svcCtx)
	return l.CreateOrderCheck(in)
}

// 回滚订单
func (s *OrderServer) RollbackOrder(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderRsp, error) {
	l := logic.NewRollbackOrderLogic(ctx, s.svcCtx)
	return l.RollbackOrder(in)
}

// 创建订单try
func (s *OrderServer) CreateOrderDTM(ctx context.Context, in *order.AddOrderReq) (*order.AddOrderRsp, error) {
	l := logic.NewCreateOrderDTMLogic(ctx, s.svcCtx)
	return l.CreateOrderDTM(in)
}

// 回撤
func (s *OrderServer) CreateOrderDTMRevert(ctx context.Context, in *order.AddOrderReq) (*order.AddOrderRsp, error) {
	l := logic.NewCreateOrderDTMRevertLogic(ctx, s.svcCtx)
	return l.CreateOrderDTMRevert(in)
}

func (s *OrderServer) GetOrderByOrdernum(ctx context.Context, in *order.GetOrderByOrdernumReq) (*order.GetOrderByOrdernumRsp, error) {
	l := logic.NewGetOrderByOrdernumLogic(ctx, s.svcCtx)
	return l.GetOrderByOrdernum(in)
}
