package model

import (
	"database/sql"
	"fmt"
	"hnchain/common/distributedid"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrdersModel = (*customOrdersModel)(nil)

type (
	// OrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrdersModel.
	OrdersModel interface {
		ordersModel
		TxUpdate(tx *sql.Tx, data *Orders) error
		TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error)
	}

	customOrdersModel struct {
		*defaultOrdersModel
	}
)

//事务更新
func (m *customOrdersModel) TxUpdate(tx *sql.Tx, data *Orders) error {

	productIdKey := fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, data.Id)
	publicOrdersOrdernumKey := fmt.Sprintf("%s%v", cachePublicOrdersOrdernumPrefix, data.Ordernum)
	
	_, err := m.Exec(func(conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set %s where id=$1",m.table,orderitemRowsWithPlaceHolder)
		return tx.Exec(query,data.Id)
	},productIdKey,publicOrdersOrdernumKey)
	if err != nil {
		return err
	}
	return nil
}

//事务插入
func(m *customOrdersModel) TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error){
	
	idgenerator := distributedid.NewSnowflake(int64(1))
	data.Id = idgenerator.GenerateId()

	publicOrdersIdKey := fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, data.Id)
	publicOrdersOrdernumKey := fmt.Sprintf("%s%v", cachePublicOrdersOrdernumPrefix, data.Ordernum)

	return m.Exec(func(conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", m.table, ordersRowsExpectAutoSet)
		return  tx.Exec(query, data.Id, data.Ordernum, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status, data.Createtime, data.Updatetime)
	},publicOrdersIdKey,publicOrdersOrdernumKey)
}

// NewOrdersModel returns a model for the database table.
func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrdersModel {
	return &customOrdersModel{
		defaultOrdersModel: newOrdersModel(conn, c, opts...),
	}
}
