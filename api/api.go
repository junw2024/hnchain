package main

import (
	"flag"
	"fmt"

	"hnchain/api/internal/config"
	"hnchain/api/internal/handler"
	"hnchain/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")
func init() {
	//close statis log
	logx.DisableStat()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
   
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//上下文环境
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
