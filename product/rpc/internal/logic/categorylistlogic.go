package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据类父ID,查询分类列表
func (l *CategoryListLogic) CategoryList(in *product.CategoryItemReq) (*product.CategoryItemRsp, error) {
	pos, err := l.svcCtx.ProductCategoryModel.ListCategoryByParentid(l.ctx,in.Parentid)
	if err != nil {
		l.Logger.Error("ListCategoryByParentid err:%v",err)
		return nil,errors.Wrap(xerr.NewErrCode(xerr.DbError),"ListCategoryByParentid err")
	}

	var categorys []*product.CategoryItem
	for _, p := range pos {
		categorys =append(categorys, &product.CategoryItem{
			Categoryid: p.Id,
			Name: p.Name,
		})
	}
	
	return &product.CategoryItemRsp{Categorys: categorys}, nil
}
