package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductCategoryModel = (*customProductCategoryModel)(nil)

type (
	// ProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCategoryModel.
	ProductCategoryModel interface {
		productCategoryModel
		ListCategoryByParentid(ctx context.Context, parentid int64)([]*ProductCategory,error)
	}

	customProductCategoryModel struct {
		*defaultProductCategoryModel
	}
)

func (m *customProductCategoryModel) ListCategoryByParentid(ctx context.Context, parentid int64)([]*ProductCategory,error){
	var categorys []*ProductCategory
	query := fmt.Sprintf("select %s from %s where parentid=$1",productCategoryRows,m.table)
	
	err :=m.QueryRowsNoCacheCtx(ctx,&categorys,query,parentid)
	if err != nil {
		return nil,err
	}
	return categorys,nil
}


// NewProductCategoryModel returns a model for the database table.
func NewProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductCategoryModel {
	return &customProductCategoryModel{
		defaultProductCategoryModel: newProductCategoryModel(conn, c, opts...),
	}
}
