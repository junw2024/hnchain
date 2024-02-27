package model

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderitemModel = (*customOrderitemModel)(nil)

type (
	// OrderitemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderitemModel.
	OrderitemModel interface {
		orderitemModel
		UserOrderitems(ctx context.Context,userid int64,ctime time.Time,size int) ([]*Orderitem,error)
	}

	customOrderitemModel struct {
		*defaultOrderitemModel
	}
)
// NewOrderitemModel returns a model for the database table.
func NewOrderitemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderitemModel {
	return &customOrderitemModel{
		defaultOrderitemModel: newOrderitemModel(conn, c, opts...),
	}
}

func (m *customOrderitemModel) UserOrderitems(ctx context.Context,userid int64,ctime time.Time,size int) ([]*Orderitem,error){
	var orderitems []*Orderitem
	query := fmt.Sprintf("select %s from %s where userid=$1 and createtime < $2 order by createtime desc  limit $3",orderitemRows,m.table)
	err := m.QueryRowsNoCacheCtx(ctx,&orderitems,query,userid,ctime,size) 
	if err != nil && err != sqlx.ErrNotFound {
		return nil,err
	}
	return orderitems,nil
}
