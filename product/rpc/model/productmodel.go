package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		CategoryProducts(ctx context.Context, ctime int64, cateid, limit int64) ([]*Product, error)
		TxUpdateStock(tx *sql.Tx, id int64, delta int) (sql.Result, error)
		UpdateProductStock(ctx context.Context, pid int64, num int32) error
	}

	customProductModel struct {
		*defaultProductModel
	}
)

func (m *customProductModel) CategoryProducts(ctx context.Context, ctime int64, cateid, limit int64) ([]*Product, error) {
	var products []*Product
	query := fmt.Sprintf("select %s from %s where categoryid=$1 and status=1 and createtime< $2 order by createtime desc limit $3",
		productRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &products, query, cateid, time.Unix(ctime, 0),limit)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	return products, nil
}

func (m *customProductModel) TxUpdateStock(tx *sql.Tx, id int64, delta int) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cachePublicProductIdPrefix, id)
	return m.Exec(func(conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set stock=stock + $1 where stock >= - $2 and id = $3", m.table)
		return tx.Exec(query, delta, delta, id)
	}, productIdKey)
}

func (m *customProductModel) UpdateProductStock(ctx context.Context, pid int64, num int32) error {
	productIdKey := fmt.Sprintf("%s%v", cachePublicProductIdPrefix, pid)
	_,err :=m.ExecCtx(ctx,func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("UPDATE %s SET stock = stock - $1 WHERE id = $2 and stock > 0",m.table)
		return conn.ExecCtx(ctx,query,num,pid)
	},productIdKey)
	return err
}

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}
