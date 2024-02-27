package logic

import (
	"context"
	"strconv"
	"strings"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/model"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductReq) (*product.ProductRsp, error) {

	products := make(map[int64]*product.ProductItem)

	ids := strings.Split(in.ProductIds,",")

	ps,err := mr.MapReduce(func(source chan<- any) {
		for _,id := range ids {
			source <- id
		}
	},func(item any, writer mr.Writer[any], cancel func(error)) {
		idstr := item.(string)
		id,err := strconv.ParseInt(idstr,10,64)
		if err != nil {
			return
		}
		p,err := l.svcCtx.ProductModel.FindOne(l.ctx,id)
		if err != nil{
			return
		}
		writer.Write(p)
	},func(pipe <- chan any, writer mr.Writer[any], cancel func(error)){
		var r[]*model.Product
		for p := range pipe {
			r = append(r,p.(*model.Product) )
			writer.Write(r)
		}
	})

	if err != nil {
		l.Logger.Errorf("ids=%s MapReduce err:%v",in.ProductIds,err)
		return nil,errors.Wrapf(xerr.NewErrCode(xerr.DbError),"%s","MapReduce query db error")
	}
	for _,p := range ps.([]*model.Product) {
		products[p.Id] = &product.ProductItem{
			Productid: p.Id,
			Name: p.Name,
			Stock: p.Stock,
		}
	}

	return &product.ProductRsp{Products: products}, nil
}
