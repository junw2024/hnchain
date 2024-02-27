package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/pkg/errors"

)

type CheckProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductStockLogic {
	return &CheckProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckProductStockLogic) CheckProductStock(in *product.UpdateProductStockReq) (*product.UpdateProductStockRsp, error) {
	p, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Productid)
	if err != nil {
		l.Logger.Error("query db error:%v",err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError),"db err!")
	}

	if p.Stock < in.Num {
		l.Logger.Info("err stock:%d,num:%d",p.Stock,in.Num)
		return nil,errors.Wrapf(xerr.NewErrCode(xerr.StockInsufficient),"stock:%d,num:%d",p.Stock,in.Num)
	}
	return &product.UpdateProductStockRsp{}, nil
}
