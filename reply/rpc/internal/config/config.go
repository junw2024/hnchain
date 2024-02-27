package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Postgres struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}
	CacheRedis cache.CacheConf
	BizRedis   redis.RedisConf
}
