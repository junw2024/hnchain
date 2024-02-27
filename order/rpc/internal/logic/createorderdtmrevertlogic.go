package logic

import (
	"context"
	"database/sql"
	"time"

	"hnchain/common/xerr"
	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/model"
	"hnchain/order/rpc/order"
	"hnchain/user/rpc/userclient"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type CreateOrderDTMRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMRevertLogic {
	return &CreateOrderDTMRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 回撤
func (l *CreateOrderDTMRevertLogic) CreateOrderDTMRevert(in *order.AddOrderReq) (*order.AddOrderRsp, error) {

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
		// 查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx,&userclient.UserInfoRequest{
			Id: in.Userid,
		})
		if err != nil {
			l.Logger.Errorf("user=%d not exixt err:%v",in.Userid,err)
			//return errors.Wrapf(xerr.NewErrMsg("user not exixt"),"msg:%v",err)
		}
		//订单
		orders,err := l.svcCtx.OrdersModel.FindOneByOrdernum(l.ctx,in.Ordernum)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("ordernum=%s not exixt err:%v",in.Ordernum,err)
			return errors.Wrapf(xerr.NewErrMsg("ordernum query err!"),"ordernum=%s",in.Ordernum)
		}
		if err == model.ErrNotFound {
			return nil
		}
		// 修改订单状态60，标识订单 已关闭
		orders.Status = 60
		orders.Updatetime = time.Now()
		return l.svcCtx.OrdersModel.TxUpdate(tx,orders)
	})

	if err != nil {
		return nil,errors.Wrap(xerr.NewErrCode(xerr.OrderRevertError),"CreateOrderDTMRevert Fail")
	}

	return &order.AddOrderRsp{Ordernum: in.Ordernum}, nil
}
