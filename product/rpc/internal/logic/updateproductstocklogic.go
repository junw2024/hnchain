package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStockLogic {
	return &UpdateProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStockLogic) UpdateProductStock(in *product.UpdateProductStockReq) (*product.UpdateProductStockRsp, error) {

	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx, in.Productid, in.Num)
	if err != nil {
		l.Logger.Error("UpdateProductStock error:%v",err)
		return nil,errors.Wrap(xerr.NewErrCode(xerr.StockInsufficient),"UpdateProductStock error")
	}
	return &product.UpdateProductStockRsp{}, nil
}
