// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productRecommendFieldNames          = builder.RawFieldNames(&ProductRecommend{}, true)
	productRecommendRows                = strings.Join(productRecommendFieldNames, ",")
	productRecommendRowsExpectAutoSet   = strings.Join(stringx.Remove(productRecommendFieldNames, "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	productRecommendRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(productRecommendFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicProductRecommendIdPrefix = "cache:public:productRecommend:id:"
)

type (
	productRecommendModel interface {
		Insert(ctx context.Context, data *ProductRecommend) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductRecommend, error)
		Update(ctx context.Context, data *ProductRecommend) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductRecommendModel struct {
		sqlc.CachedConn
		table string
	}

	ProductRecommend struct {
		Id         int64     `db:"id"`
		Productid  int64     `db:"productid"`
		Status     int32     `db:"status"`   // 状态1:在推荐 0:不推荐
		Heat       int32     `db:"heat"`     // 推荐指数，越大越靠前
		Imageurl   string    `db:"imageurl"` // 推荐主图
		Createtime time.Time `db:"createtime"`
		Updatetime time.Time `db:"updatetime"`
	}
)

func newProductRecommendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductRecommendModel {
	return &defaultProductRecommendModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."product_recommend"`,
	}
}

func (m *defaultProductRecommendModel) Delete(ctx context.Context, id int64) error {
	publicProductRecommendIdKey := fmt.Sprintf("%s%v", cachePublicProductRecommendIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicProductRecommendIdKey)
	return err
}

func (m *defaultProductRecommendModel) FindOne(ctx context.Context, id int64) (*ProductRecommend, error) {
	publicProductRecommendIdKey := fmt.Sprintf("%s%v", cachePublicProductRecommendIdPrefix, id)
	var resp ProductRecommend
	err := m.QueryRowCtx(ctx, &resp, publicProductRecommendIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", productRecommendRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductRecommendModel) Insert(ctx context.Context, data *ProductRecommend) (sql.Result, error) {
	publicProductRecommendIdKey := fmt.Sprintf("%s%v", cachePublicProductRecommendIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7)", m.table, productRecommendRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Productid, data.Status, data.Heat, data.Imageurl, data.Createtime, data.Updatetime)
	}, publicProductRecommendIdKey)
	return ret, err
}

func (m *defaultProductRecommendModel) Update(ctx context.Context, data *ProductRecommend) error {
	publicProductRecommendIdKey := fmt.Sprintf("%s%v", cachePublicProductRecommendIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, productRecommendRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Productid, data.Status, data.Heat, data.Imageurl, data.Createtime, data.Updatetime)
	}, publicProductRecommendIdKey)
	return err
}

func (m *defaultProductRecommendModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicProductRecommendIdPrefix, primary)
}

func (m *defaultProductRecommendModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", productRecommendRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductRecommendModel) tableName() string {
	return m.table
}
