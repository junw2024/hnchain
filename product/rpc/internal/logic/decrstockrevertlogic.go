package logic

import (
	"context"
	"database/sql"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type DecrStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockRevertLogic {
	return &DecrStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockRevertLogic) DecrStockRevert(in *product.DecrStockReq) (*product.DecrStockRsp, error) {
	//db
	db, err := postgres.New(l.svcCtx.Config.Postgres.DataSource).RawDB()
	if err != nil {
		l.Logger.Errorf("get DB error:%v", err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "get DB error!")
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)

	if err != nil {
		l.Logger.Errorf("BarrierFromGrpc err:%v", err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.TransactionError), "BarrierFromGrpc err!")
	}

	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		var num int =int(in.Num)
		//更新产品库存
		_, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.Id, num)
		return err
	})
	if err != nil {
		return nil ,errors.Wrap(xerr.NewErrCode(xerr.StockInsufficient),"reback stock err!")
	}
	return &product.DecrStockRsp{}, nil
}
