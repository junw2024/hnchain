package user

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionDelLogic {
	return &UserCollectionDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionDelLogic) UserCollectionDel(req *types.UserCollectionDelReq) (resp *types.UserCollectionDelRes, err error) {
    var rpcReq userclient.UserCollectionDelReq
	rpcReq.Id=req.Id
	_,err = l.svcCtx.UserRPC.DelUserCollection(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to UserCollectionAdd err : %v ,req:%+v", err, req)
		return nil,err
	}
	return &types.UserCollectionDelRes{},nil
}
