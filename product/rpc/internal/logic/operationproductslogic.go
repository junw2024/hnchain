package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type OperationProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	productListLogic *ProductListLogic
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		productListLogic: NewProductListLogic(ctx, svcCtx),
	}
}

const (
	status               = 1
	operationProductsKey = "operation#products"
)

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductsReq) (*product.OperationProductsRsp, error) {
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)
	if ok {
		return &product.OperationProductsRsp{Products: opProducts.([]*product.ProductItem)}, nil
	}

	pos, err := l.svcCtx.ProductOperationModel.OperationProducts(l.ctx, status)
	if err != nil {
		l.Logger.Errorf("query ProdcutOperation err:%v", err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "query ProdcutOperation err:")
	}
	var ips []int64
	for _, p := range pos {
		ips = append(ips, p.Productid)
	}

	products, err := l.productListLogic.productsByIds(l.ctx, ips)
	if err != nil {
		l.Logger.Errorf("products By Ids:%v", err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "query products err:")
	}

	var pItems []*product.ProductItem

	for _,p := range products {
		pItems = append(pItems, &product.ProductItem{
			Productid: p.Id,
			Name: p.Name,
		})
	}
	l.svcCtx.LocalCache.Set(operationProductsKey,pItems)
	return &product.OperationProductsRsp{Products: pItems}, nil
}
