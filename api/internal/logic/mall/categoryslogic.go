package mall

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/product/rpc/productclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategorysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategorysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategorysLogic {
	return &CategorysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategorysLogic) Categorys(req *types.CategoryReq) (resp *types.CategoryRsp, err error) {

	var categoryItemReq productclient.CategoryItemReq
	categoryItemReq.Parentid = req.Parentid
	ps,err := l.svcCtx.ProductRPC.CategoryList(l.ctx,&categoryItemReq)

	if err != nil {
		l.Logger.Errorf("Categorys err:%v",err)
		return nil,err
	}
	
	var categorys []*types.Category
    for _,p := range ps.Categorys {
		categorys =append(categorys, &types.Category{
			Id: p.Categoryid,
			Name: p.Name,
		})
	}
	return &types.CategoryRsp{Categorys: categorys},nil
}
