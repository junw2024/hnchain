package logic

import (
	"context"
	"fmt"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/model"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemReq) (*product.ProductItem, error) {
	//并发合并请求:N请求合并1一个请求，这样减少数据库压力
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product:%d", in.Productid), func() (interface{}, error) {
		return l.svcCtx.ProductModel.FindOne(l.ctx, in.Productid)
	})

	if err != nil {
		l.Logger.Errorf("Find id=%d data err:%v",in.Productid,err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ProductExistError),"product error:%v",err)
	}
	p := v.(*model.Product)
	return &product.ProductItem{
		Productid: p.Id,
		Name: p.Name,
		Stock: p.Stock,
		Description: p.Detail,
		Imageurl: p.Imageurl,
		
	}, nil
}
