package user

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/xerr"
	"hnchain/user/rpc/userclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelRevAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelRevAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRevAddressLogic {
	return &DelRevAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRevAddressLogic) DelRevAddress(req *types.UserRevAddrDelReq) ( *types.UserRevAddrDelRes, error) {
	var rpcReq userclient.UserRevAddrDelReq
	rpcReq.Id= req.Id
	_,err := l.svcCtx.UserRPC.DelUserRevAddr(l.ctx,&rpcReq)
	if err != nil {
		return nil,errors.Wrapf(xerr.NewErrMsg("DelRevAddress error!"), "req: %+v", req)
	}
	return &types.UserRevAddrDelRes{},nil
}
