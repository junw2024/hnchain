package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HnuserModel = (*customHnuserModel)(nil)

type (
	// HnuserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHnuserModel.
	HnuserModel interface {
		hnuserModel
	}

	customHnuserModel struct {
		*defaultHnuserModel
	}
)

// NewHnuserModel returns a model for the database table.
func NewHnuserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HnuserModel {
	
	return &customHnuserModel{
		defaultHnuserModel: newHnuserModel(conn, c, opts...),
	}
	
}
