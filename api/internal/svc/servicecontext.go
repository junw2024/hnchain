package svc

import (
	"hnchain/api/internal/config"
	"hnchain/order/rpc/orderclient"
	"hnchain/product/rpc/productclient"
	"hnchain/reply/rpc/replyclient"
	"hnchain/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	UserRPC  userclient.User
	ProductRPC productclient.Product
	OrderRPC  orderclient.Order
	ReplyRPC replyclient.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRPC: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ReplyRPC: replyclient.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
	}
}
