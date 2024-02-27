package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductOperationModel = (*customProductOperationModel)(nil)

type (
	// ProductOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductOperationModel.
	ProductOperationModel interface {
		productOperationModel
		OperationProducts(ctx context.Context, status int32) ([]*ProductOperation, error)
	}

	customProductOperationModel struct {
		*defaultProductOperationModel
	}
)

// NewProductOperationModel returns a model for the database table.
func NewProductOperationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductOperationModel {
	return &customProductOperationModel{
		defaultProductOperationModel: newProductOperationModel(conn, c, opts...),
	}
}

func (m *customProductOperationModel) OperationProducts(ctx context.Context, status int32) ([]*ProductOperation, error) {
	var operations []*ProductOperation
	err := m.QueryRowsNoCacheCtx(ctx,
		&operations,
		fmt.Sprintf("select %s from %s where status=$1",productOperationRows,m.table),
		 status)
   	if err != nil {
		return nil,err
	}
	return operations,nil
}
