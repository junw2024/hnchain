package mall

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/product/rpc/productclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.CategoryListReq) (resp *types.CategoryListRsp, err error) {
	
	l.Logger.Errorv(req)

	var rpcReq productclient.ProductListReq
	rpcReq.Categoryid = req.Categoryid
	rpcReq.Cursor = req.Cursor
	rpcReq.Ps = req.Ps
	rpcReq.Productid = req.Productid

	ps,err :=l.svcCtx.ProductRPC.ProductList(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("CategoryList err:%v",err)
		return nil,err
	}

	var products []*types.Product
	for _, p := range ps.Products {

		var it types.Product
		_=copier.Copy(&it,p)
		it.Id = p.Productid
		
		products = append(products,&it)
	}
	
	return  &types.CategoryListRsp{
		Products: products,
		IsEnd:    ps.IsEnd,
		Productid: ps.Productid,
		LastVal:  ps.Timestamp,
	},nil
}
