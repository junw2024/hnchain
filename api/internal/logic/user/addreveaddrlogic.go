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

type AddReveAddrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReveAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReveAddrLogic {
	return &AddReveAddrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReveAddrLogic) AddReveAddr(req *types.UserRevAddrAddReq) (*types.UserRevAddrAddRes, error) {
	var rpcReq userclient.UserRevAddrAddReq
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		l.Logger.Errorf("token:Failed to AddReveAddr err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("AddReveAddr error! Get token"),"")
	}
	
	rpcReq.Uid = uid
	_ = copier.Copy(&rpcReq,req)

	_,err = l.svcCtx.UserRPC.AddUserRevAddr(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to AddReveAddr err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("post AddUserRevAddr error!"),"rpc AddUserRevAddr error!")
	}
	return &types.UserRevAddrAddRes{},nil
}
