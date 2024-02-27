package logic

import (
	"context"
	"fmt"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndUpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStockLogic {
	return &CheckAndUpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	luaCheckAndUpdateScript = `
local counts = redis.call("HMGET",KEYS[1],"total", "seckill")
local total = tonumber(counts[1])
local seckill = tonumber(seckill)
if seckill + 1  <= count then
	redis.call("HINCRBY", KEYS[1], "seckill", 1)
	return 1
end 
return 0
`
)

func (l *CheckAndUpdateStockLogic) CheckAndUpdateStock(in *product.CheckAndUpdateStockReq) (*product.CheckAndUpdateStockRsp, error) {
	val, err := l.svcCtx.BizRedis.EvalCtx(l.ctx,luaCheckAndUpdateScript,[]string{stockKey(in.Productid)})
	if err != nil {
		l.Logger.Errorf("BizRedis stock error:%v",err)
		return nil,errors.Wrapf(xerr.NewErrCode(xerr.RedisError),"stock扣减异常")
	}
	if val.(int64) == 0 {
		l.Logger.Info("BizRedis stock insufficient:productid:%d ",in.Productid)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.StockInsufficient),"insufficient stock: %d",in.Productid)
	}

	return &product.CheckAndUpdateStockRsp{}, nil
}

func stockKey(pid int64)  string{
	return fmt.Sprintf("stock%d",pid)
}
