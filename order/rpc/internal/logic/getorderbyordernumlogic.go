package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/model"
	"hnchain/order/rpc/order"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByOrdernumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByOrdernumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByOrdernumLogic {
	return &GetOrderByOrdernumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByOrdernumLogic) GetOrderByOrdernum(in *order.GetOrderByOrdernumReq) (*order.GetOrderByOrdernumRsp, error) {
	p,err := l.svcCtx.OrdersModel.FindOneByOrdernum(l.ctx,in.Ordernum)
	if err != nil && err == model.ErrNotFound {
		l.Logger.Errorf("ordernum=%s is not found",in.Ordernum)
		return nil,errors.Wrapf(xerr.NewErrCode(xerr.DataNoExistError),"ordernum:%s not found",in.Ordernum)
	}
	if err != nil {
		l.Logger.Errorf("GetOrderByOrdernum err:%v",err)
		return nil,errors.Wrap(xerr.NewErrCode(xerr.DbError),"db query err!")
	}

	var orders order.Orders
	copier.Copy(&orders,p)
	orders.CreateTime = p.Createtime.Unix()
	orders.UpdateTime = p.Updatetime.Unix()

	return &order.GetOrderByOrdernumRsp{Orders: &orders}, nil
}
