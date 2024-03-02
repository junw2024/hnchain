package order

import (
	"context"
	"encoding/json"
	"fmt"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/distributedid"
	"hnchain/common/xerr"
	"hnchain/order/rpc/orderclient"
	"hnchain/product/rpc/productclient"
	"time"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 分布式事务
func (l *AddOrderLogic) AddOrder(req *types.OrderAddReq) (resp *types.OrderAddRsp, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	// 获取 OrderRpc BuildTarget
	orderRpcServer, err := l.svcCtx.Config.OrderRPC.BuildTarget()
	if err != nil {
		return nil, errors.Wrap(xerr.NewErrMsg("create order excpetion"), "AddOrder orderRpcServer")
	}
	// 获取 ProductRpc BuildTarget
	productRpcServer, err := l.svcCtx.Config.ProductRPC.BuildTarget()
	if err != nil {
		return nil, errors.Wrap(xerr.NewErrMsg("create order excpetion"), "AddOrder productRpcServer")
	}

	var ordernum = genOrdernum()
	// dtm 服务的 etcd 注册地址
	var dtmserver = "discov://127.0.0.1:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmserver)
	// 创建一个saga协议的事务
	var rpcOrderAddReq = orderclient.AddOrderReq{
		Userid:        uid,
		Ordernum:      ordernum,
		Productid:     req.Productid,
		Quantity:      req.Quantity,
		Postage:       req.Postage,
		ReceiveAddrId: req.RevAddrId,
	}

	// 创建一个saga协议的事务
   saga := dtmgrpc.NewSagaGrpc(dtmserver,gid).
   Add(
	orderRpcServer+"/orderclient.Order/CreateOrderDTM",
	orderRpcServer+"/orderclient.Order/CreateOrderDTMRevert",
	&rpcOrderAddReq).
	Add(
		productRpcServer+"productclient.Product/DecrStock",
		productRpcServer+"productRpcServer+productclient.Product.DecrStockRevert",
		&productclient.DecrStockReq{
		Id: req.Productid,
		Num: req.Quantity,
	})
	// 事务提交
	err = saga.Submit()
    if err != nil {
		l.Logger.Infof("订单saga事务异常:%+v",err)
		return nil, errors.Wrap(xerr.NewErrMsg("create order excpetion"), "AddOrder saga tx")
	}

	return &types.OrderAddRsp{Ordernum: ordernum},nil
}

// 生产订单号
func genOrdernum() string {
	idgenerator := distributedid.NewSnowflake(int64(1))
	return fmt.Sprintf("%s-%d", time.Now().Format("20060102"), idgenerator.GenerateId())
}
