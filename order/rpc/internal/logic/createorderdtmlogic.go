package logic

import (
	"context"
	"database/sql"
	"hnchain/common/xerr"
	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/model"
	"hnchain/order/rpc/order"
	"hnchain/product/rpc/productclient"
	"hnchain/user/rpc/userclient"
	"time"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/shopspring/decimal"
)

type CreateOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMLogic {
	return &CreateOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建订单try
func (l *CreateOrderDTMLogic) CreateOrderDTM(in *order.AddOrderReq) (*order.AddOrderRsp, error) {
	var (
		userRpcRs    *userclient.UserInfoResponse
		productRpcRs *productclient.ProductItem
		receiveAddr  *userclient.UserRevAddr
	)

	//check product
	checkProduct := func() error {
		var err error
		var productReq productclient.ProductItemReq
		productReq.Productid = in.Productid
		productRpcRs, err = l.svcCtx.ProductRpc.Product(l.ctx, &productReq)
		if err != nil {
			l.Logger.Errorf("checkProduct err:%v", err)
			return err
		}
		return nil
	}

	//check user
	checkUser := func() error {
		var err error
		var userReq userclient.UserInfoRequest
		userReq.Id = in.Userid
		userRpcRs, err = l.svcCtx.UserRpc.UserInfo(l.ctx, &userReq)
		if err != nil {
			l.Logger.Errorf("checkUser err:%v", err)
			return err
		}
		return nil
	}
	//check userRevAddr
	ckeckRecAddr := func() error {
		var err error
		var receiveAddrReq userclient.UserRevAddrInfoReq
		receiveAddrReq.Id = in.ReceiveAddrId
		receiveAddr, err = l.svcCtx.UserRpc.GetUserRevAddrInfo(l.ctx, &receiveAddrReq)
		if err != nil {
			return err
		}
		return nil
	}

	//parallel call:并发调用
	err := mr.Finish(checkUser, checkProduct, ckeckRecAddr)

	if userRpcRs == nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DataNoExistError), "error! user not exist exception")
	}
	if productRpcRs == nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DataNoExistError), "error! product not exist exception")
	}
	//检查库存
	if productRpcRs.Stock < in.Quantity {
		return nil, errors.Wrapf(xerr.NewErrMsg("product understock"), "product understock")
	}

	if receiveAddr == nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DataNoExistError), "error! receiveAddr not exist exception")
	}

	ordernum := in.Ordernum
	//db
	db, err := postgres.New(l.svcCtx.Config.Postgres.DataSource).RawDB()
	if err != nil {
		l.Logger.Errorf("postgre db err:%v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("postgre db err!"), "msg:%v", err)
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("BarrierFromGrpc err!"), "msg:%v", err)
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		//create new shipping
		var shipping model.Shipping
		shipping.Createtime = time.Now()
		shipping.Updatetime = time.Now()
		shipping.Ordernum = ordernum
		shipping.ReceiverName = receiveAddr.Name
		shipping.ReceiverPhone = receiveAddr.Phone
		shipping.ReceiverMobile = receiveAddr.Phone
		shipping.ReceiverProvince = receiveAddr.Province
		shipping.ReceiverCity = receiveAddr.City
		shipping.ReceiverAddress = receiveAddr.DetailAddress
		shipping.Userid = in.Userid

		_,err = l.svcCtx.ShippingModel.Insert(l.ctx, &shipping)
		if err != nil {
			l.Logger.Errorf("ShippingModel.Insert err:%v", err)
			return errors.Wrapf(xerr.NewErrMsg("Shipping:insert db error!"), "msg:%v", err)
		}


		//create new orderitem
		orderitem := model.Orderitem{
			Ordernum:     ordernum,
			Userid:       in.Userid,
			Productid:    in.Productid,
			Productname:  productRpcRs.Name,
			Productimage: productRpcRs.Imageurl,
			Currentprice: productRpcRs.Price,
			Quantity: in.Quantity,
		}
		//计算金额
		totalprice :=decimal.NewFromFloat(orderitem.Currentprice).Mul(decimal.NewFromInt32(orderitem.Quantity))
		orderitem.Totalprice = totalprice.InexactFloat64()
		orderitem.Createtime = time.Now()
		orderitem.Updatetime = time.Now()
        _,err = l.svcCtx.OrderitemModel.Insert(l.ctx,&orderitem)
		if err != nil {
			l.Logger.Errorf("OrderitemModel.Insert err:%v", err)
			return errors.Wrapf(xerr.NewErrMsg("Orderitem:insert db error!"), "msg:%v", err)
		}
	
		//create new order
		insertOrders := model.Orders{
			Ordernum: ordernum,
			Userid: in.Userid,
			Shoppingid: shipping.Id,
			Paymenttype: 1,
			Postage: in.Postage,
			Status: 10,
			Createtime: time.Now(),
			Updatetime: time.Now(),
		}
		//计算金额
        amount := decimal.NewFromFloat(orderitem.Totalprice).Add(decimal.NewFromFloat(in.Postage))
		insertOrders.Payment = amount.InexactFloat64()

		//执行事务
		_,err = l.svcCtx.OrdersModel.TxInsert(tx,&insertOrders)
		if err != nil {
			l.Logger.Errorf("OrdersModel.Insert err:%v", err)
			return errors.Wrapf(xerr.NewErrMsg("Orders:insert db error!"), "msg:%v", err)
		}
		return nil
	})

	if err != nil {
		return nil,errors.Wrap(xerr.NewErrCode(xerr.OrderCreateError),"create order error!")
	}

	return &order.AddOrderRsp{Ordernum: ordernum}, nil
}


