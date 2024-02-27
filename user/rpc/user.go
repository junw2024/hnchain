package main

import (
	"flag"
	"fmt"

	"hnchain/user/rpc/internal/config"
	"hnchain/user/rpc/internal/server"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()
     
	//close statis log
	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)


	//上下文环境
	ctx := svc.NewServiceContext(c)
    //service rpc
	svr := server.NewUserServer(ctx)

	//启动服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		//注册服务
		user.RegisterUserServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
