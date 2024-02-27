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

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockLogic) DecrStock(in *product.DecrStockReq) (*product.DecrStockRsp, error) {
	//db
	db, err := postgres.New(l.svcCtx.Config.Postgres.DataSource).RawDB()
	if err != nil {
		l.Logger.Errorf("get DB error:%v",err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError),"get DB error!")
	}

	// 获取子事务屏障对象 ctx通过上下文传递
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Logger.Errorf("BarrierFromGrpc err:%v",err)
		return nil, errors.Wrap(xerr.NewErrCode(xerr.TransactionError),"BarrierFromGrpc err!")
	} 

	// 开启子事务屏障
    err = barrier.CallWithDB(db,func(tx *sql.Tx) error {
		//更新产品库存
		result, err :=  l.svcCtx.ProductModel.TxUpdateStock(tx,in.Id,-1)
		if err != nil {
			l.Logger.Errorf("udpate product stock err!:%v",err)
			return xerr.NewErrCode(xerr.TransactionError)
		}

		affected, err := result.RowsAffected() 
		if err == nil && affected == 0 {
			l.Logger.Error("product stock StockInsufficient")
			return xerr.NewErrCode(xerr.StockInsufficient)
		}
		
		return nil
	})
	// 库存不足，不再重试，走回滚
	if err != nil && err.(* xerr.CodeError).GetErrCode() == xerr.StockInsufficient  {
		return nil,errors.Wrap(err,"stock StockInsufficient")
	}
	if err != nil {
		return nil,errors.Wrap(err,"udpate product stock  err!")
	}
	return &product.DecrStockRsp{}, nil
}