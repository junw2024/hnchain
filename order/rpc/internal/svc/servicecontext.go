package svc

import (
	"hnchain/order/rpc/internal/config"
	"hnchain/order/rpc/model"
	"hnchain/product/rpc/productclient"
	"hnchain/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrdersModel model.OrdersModel
	OrderitemModel model.OrderitemModel
	ShippingModel model.ShippingModel
	UserRpc userclient.User
	ProductRpc productclient.Product
	BizRedis              *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := postgres.New(c.Postgres.DataSource)
	db, err := sqlconn.RawDB()
	if err == nil {
		db.SetMaxIdleConns(c.Postgres.MaxIdleConns)
		db.SetMaxOpenConns(c.Postgres.MaxOpenConns)
	}

	return &ServiceContext{
		Config: c,
		OrdersModel: model.NewOrdersModel(sqlconn,c.CacheRedis),
		OrderitemModel: model.NewOrderitemModel(sqlconn,c.CacheRedis),
		ShippingModel: model.NewShippingModel(sqlconn,c.CacheRedis),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		BizRedis: redis.New(c.BizRedis.Host,redis.WithPass(c.BizRedis.Pass)),
	}
}
