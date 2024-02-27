package user

import (
	"context"
	"encoding/json"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/xerr"
	"hnchain/user/rpc/userclient"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRevAddrListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRevAddrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRevAddrListLogic {
	return &UserRevAddrListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRevAddrListLogic) UserRevAddrList(req *types.UserRevAddrListReq) (*types.UserRevAddrListRes,  error) {
	var rpcReq userclient.UserRevAddrListReq
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
    if err != nil {
		l.Logger.Errorf("token:Failed to get user addrerss  list err : %v ,req:%+v", err, req)
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! get uid from token"),"")
	}
	
	rpcReq.Uid= uid

	rpcRes,err := l.svcCtx.UserRPC.GetUserRevAddrList(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to get user addrerss  list err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("Error! Function GetUserRevAddrList"),"error! rpc GetUserRevAddrList")
	}
	var list []types.UserRevAddr
	for _, it := range rpcRes.List {
		var revAddrVo types.UserRevAddr
		_ = copier.Copy(&revAddrVo,it)
		list=append(list, revAddrVo)
	}
	return &types.UserRevAddrListRes{List: list},nil
}
