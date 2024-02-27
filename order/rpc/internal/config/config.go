package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/core/stores/redis"

)

type Config struct {
	zrpc.RpcServerConf
	Postgres struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}
	CacheRedis cache.CacheConf
	ProductRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	BizRedis   redis.RedisConf
}
