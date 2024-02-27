package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HnuserCollectionModel = (*customHnuserCollectionModel)(nil)

type (
	// HnuserCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHnuserCollectionModel.
	HnuserCollectionModel interface {
		hnuserCollectionModel
		UpdateIsDelete(ctx context.Context, data *HnuserCollection) error
		FindAllByUid(ctx context.Context, uid int64) ([]*HnuserCollection, error)
	}

	customHnuserCollectionModel struct {
		*defaultHnuserCollectionModel
	}
)

// NewHnuserCollectionModel returns a model for the database table.
func NewHnuserCollectionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HnuserCollectionModel {
	return &customHnuserCollectionModel{
		defaultHnuserCollectionModel: newHnuserCollectionModel(conn, c, opts...),
	}
}

func (m *customHnuserCollectionModel) UpdateIsDelete(ctx context.Context, data *HnuserCollection) error {
	userCollectionIdKey := fmt.Sprintf("%s%v", cachePublicHnuserCollectionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set isdelete = true,updatetime=$2 where id=$1", m.table)
		return conn.ExecCtx(ctx, query, data.Id, data.Updatetime)

	}, userCollectionIdKey)

	return err
}

func (m *customHnuserCollectionModel) FindAllByUid(ctx context.Context, uid int64) ([]*HnuserCollection, error) {
	var resp []*HnuserCollection
	query := fmt.Sprintf("select %s from %s where uid=$1 and isdelete=false", hnuserCollectionRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil,err
	}
}
