package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductRecommendModel = (*customProductRecommendModel)(nil)

type (
	// ProductRecommendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductRecommendModel.
	ProductRecommendModel interface {
		productRecommendModel
		ProductRecommens(ctx context.Context,num int32) ([]*ProductRecommend,error)
	}

	customProductRecommendModel struct {
		*defaultProductRecommendModel
	}
)

func(m *customProductRecommendModel)ProductRecommens(ctx context.Context,num int32) ([]*ProductRecommend,error) {
	var recommends []*ProductRecommend
	query := fmt.Sprintf("select %s from %s where status=1 order by heat desc limit $1",productRecommendRows,m.table)

	//fmt.Printf("query:%s\n",query)

	err := m.QueryRowsNoCacheCtx(ctx,&recommends,query,num)
	if err != nil {
		return nil,err
	}
	return recommends,nil
}

// NewProductRecommendModel returns a model for the database table.
func NewProductRecommendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductRecommendModel {
	return &customProductRecommendModel{
		defaultProductRecommendModel: newProductRecommendModel(conn, c, opts...),
	}
}
