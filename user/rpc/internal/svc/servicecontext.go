package svc

import (
	"hnchain/user/rpc/internal/config"
	"hnchain/user/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config config.Config
	//add dependency on user model
	HnuserModel model.HnuserModel
	HnuserCollectionModel model.HnuserCollectionModel
	HnuserRevAddrModel model.HnuserRevAddrModel
}


func NewServiceContext(c config.Config) *ServiceContext {
	
    sqlcon :=postgres.New(c.Postgres.DataSource)
	db,err  :=sqlcon.RawDB()
	if err == nil {
		//db.SetConnMaxLifetime(time.Second)
		db.SetMaxIdleConns(c.Postgres.MaxIdleConns)
		db.SetMaxOpenConns(c.Postgres.MaxOpenConns)
	}
	return &ServiceContext{
		Config: c,
		HnuserModel: model.NewHnuserModel(sqlcon,c.CacheRedis),
		HnuserCollectionModel: model.NewHnuserCollectionModel(sqlcon,c.CacheRedis),
		HnuserRevAddrModel: model.NewHnuserRevAddrModel(sqlcon,c.CacheRedis),
	}
}
