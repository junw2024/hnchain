package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/model"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductRecommendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductRecommendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductRecommendsLogic {
	return &ProductRecommendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 目标数量的推荐商品
func (l *ProductRecommendsLogic) ProductRecommends(in *product.ProductRecommendReq) (*product.ProductRecommendRsp, error) {

	rs, err := l.svcCtx.ProductRecommendModel.ProductRecommens(l.ctx,in.Num)
	if err != nil {
		l.Logger.Errorf("query productRecommend err:%s",err)
		return nil,errors.Wrap(xerr.NewErrCode(xerr.DbError),"query product recommend err!")
	}
	var products []*product.ProductItem
	for _, it := range rs {
		products =append(products, &product.ProductItem{
			Productid: it.Productid,
			Imageurl: it.Imageurl,
		})
	}

	mp,err := mr.MapReduce(func(source chan<- int64) {
		for _,it := range products {
			source <- it.Productid
		}
	},func(it int64, writer mr.Writer[*model.Product],cancal func(error)) {
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx,it)
		if err != nil {
			cancal(err)
			return
		}
		writer.Write(p)
	},func(pipe <-chan *model.Product, writer mr.Writer[map[int64]string], cancel func(error)) {
		rs := make(map[int64]string,0)
		for it := range pipe {
			rs[it.Id]=it.Name
		}
		writer.Write(rs)
	})

	if err != nil {
		return nil,errors.Wrap(xerr.NewErrCode(xerr.DbError),"find product err!")
	}
	
	for k,v := range mp {
		it := find(k,products)
		if it != nil {
			it.Name=v
		}
	}

	return &product.ProductRecommendRsp{Products: products}, nil
}

func find(id int64,products []*product.ProductItem) *product.ProductItem {
	for _,it := range products {
		if it.Productid == id {
			return it
		}
	}
	return nil
}
