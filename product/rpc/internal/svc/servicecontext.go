package svc

import (
	"hnchain/product/rpc/internal/config"
	"hnchain/product/rpc/model"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"golang.org/x/sync/singleflight"
)
const localCacheExpire = time.Duration(time.Second * 60)

type ServiceContext struct {
	Config                config.Config
	ProductModel          model.ProductModel
	ProductCategoryModel  model.ProductCategoryModel
	ProductOperationModel model.ProductOperationModel
	ProductRecommendModel model.ProductRecommendModel  
	BizRedis              *redis.Redis
	LocalCache            *collection.Cache
	SingleGroup           singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlconn := postgres.New(c.Postgres.DataSource)
	db, err := sqlconn.RawDB()
	if err == nil {
		db.SetMaxIdleConns(c.Postgres.MaxIdleConns)
		db.SetMaxOpenConns(c.Postgres.MaxOpenConns)
	}

	localCache,err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	
	return &ServiceContext{
		Config:                c,
		ProductModel:          model.NewProductModel(sqlconn, c.CacheRedis),
		ProductCategoryModel:  model.NewProductCategoryModel(sqlconn, c.CacheRedis),
		ProductOperationModel: model.NewProductOperationModel(sqlconn, c.CacheRedis),
		ProductRecommendModel: model.NewProductRecommendModel(sqlconn,c.CacheRedis),
		BizRedis: redis.New(c.BizRedis.Host,redis.WithPass(c.BizRedis.Pass)),
		LocalCache: localCache,
	}
}
