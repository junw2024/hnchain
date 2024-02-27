package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HnuserRevAddrModel = (*customHnuserRevAddrModel)(nil)

type (
	// HnuserRevAddrModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHnuserRevAddrModel.
	HnuserRevAddrModel interface {
		hnuserRevAddrModel
		UpdateIsDelete(ctx context.Context, data *HnuserRevAddr) error
		FindAllByUid(ctx context.Context, uid int64) ([]*HnuserRevAddr, error)
		FindAllIdByUid(ctx context.Context,uid int64) ([]*HnuserRevAddrId,error)
		EditUserRevAddr(ctx context.Context,data *HnuserRevAddr) error
	}

	customHnuserRevAddrModel struct {
		*defaultHnuserRevAddrModel
	}
)

// NewHnuserRevAddrModel returns a model for the database table.
func NewHnuserRevAddrModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HnuserRevAddrModel {
	return &customHnuserRevAddrModel{
		defaultHnuserRevAddrModel: newHnuserRevAddrModel(conn, c, opts...),
	}
}

func (m *customHnuserRevAddrModel) UpdateIsDelete(ctx context.Context, data *HnuserRevAddr) error {
	userReceiveAddressIdKey := fmt.Sprintf("%s%v", cachePublicHnuserRevAddrIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {

		query := fmt.Sprintf("update %s set isdelete=$2,updatetime=$3 where id =$1", m.table)
		return conn.ExecCtx(ctx, query, data.Id,data.Isdelete,data.Updatetime)

	}, userReceiveAddressIdKey)
	return err
}

func (m *customHnuserRevAddrModel) FindAllByUid(ctx context.Context, uid int64) ([]*HnuserRevAddr, error) {
	var resp []*HnuserRevAddr
	query := fmt.Sprintf("select %s from %s where uid=$1 and isdelete=false", hnuserRevAddrRows, m.table)
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

func (m *customHnuserRevAddrModel) FindAllIdByUid(ctx context.Context,uid int64) ([]*HnuserRevAddrId,error) {
	ids := make([]*HnuserRevAddrId,0)
	query := fmt.Sprintf("select id from %s where uid =$1",m.table)
	err := m.QueryRowsNoCacheCtx(ctx,&ids,query,uid)
	switch err {
	case nil:
		return ids, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil,err
	}
}

func (m *customHnuserRevAddrModel) EditUserRevAddr(ctx context.Context,data *HnuserRevAddr) error {
	if data.Isdefault == bool(true) {
		ids,err := m.FindAllIdByUid(ctx,data.Uid)
		if err != nil && err != ErrNotFound {
			return err
		}
		logx.Errorf("error!:%v",ids)
		for _, it := range ids {
			publicHnuserRevAddrIdKey := fmt.Sprintf("%s%v", cachePublicHnuserRevAddrIdPrefix, data.Id)
			m.ExecCtx(ctx,func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
				query :=fmt.Sprintf("update %s set isdefault=false,updatetime=$2 where id=$1",m.table)
				return conn.ExecCtx(ctx,query,it.Id,time.Now())
			},publicHnuserRevAddrIdKey)
		}
	}
	return m.Update(ctx,data)
}
	