package mall

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/xerr"
	"hnchain/product/rpc/productclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend(req *types.RecommendReq) (resp *types.RecommendRsp, err error) {
	
	var rpcReq productclient.ProductRecommendReq
	rpcReq.Num=req.Ps
	ps,err :=l.svcCtx.ProductRPC.ProductRecommends(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("ProductRPC.ProductRecommends:%v",err)
		return nil,errors.Wrap(xerr.NewErrCode(xerr.ServerCommonError),".ProductRPC.ProductRecommends err!")
	}

	var products []*types.Product
	for _,p := range ps.Products {
		products=append(products, &types.Product{
			Id: p.Productid,
			Name: p.Name,
			Imageurl: p.Imageurl,
		})
	}
	return &types.RecommendRsp{Products: products},nil
}
