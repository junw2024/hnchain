package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackProductStockLogic {
	return &RollbackProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackProductStockLogic) RollbackProductStock(in *product.UpdateProductStockReq) (*product.UpdateProductStockRsp, error) {
	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx,in.Productid,-in.Num)
	if err != nil {
		l.Logger.Errorf("RollbackProductStock err:%v",err)
		return nil,errors.Wrap(xerr.NewErrCode(xerr.StockInsufficient),"RollbackProductStock err!")
	}
	return &product.UpdateProductStockRsp{}, nil
}
