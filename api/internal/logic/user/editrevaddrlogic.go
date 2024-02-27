package user

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/xerr"
	"hnchain/user/rpc/userclient"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditRevAddrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditRevAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditRevAddrLogic {
	return &EditRevAddrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditRevAddrLogic) EditRevAddr(req *types.UserRevAddrEditReq) (resp *types.UserRevAddrEditRes, err error) {
	var rpcReq userclient.UserRevAddrEditReq
	_ = copier.Copy(&rpcReq,req)
	
	_,err = l.svcCtx.UserRPC.EditUserRevAddr(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to EditRevAddr err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("Fail EditRevAddr"),"rpc EditUserRevAddr fail")
	}
	return &types.UserRevAddrEditRes{},nil
}
