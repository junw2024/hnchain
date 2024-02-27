package user

import (
	"context"
	"encoding/json"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/xerr"
	"hnchain/user/rpc/userclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionAddLogic {
	return &UserCollectionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionAddLogic) UserCollectionAdd(req *types.UserCollectionAddReq) (resp *types.UserCollectionAddRes, err error) {
	uid,_ := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		l.Logger.Errorf("token:Failed to UserCollectionAdd err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("UserCollectionAdd error! Get token"),"")
	}
	var rpcReq userclient.UserCollectionAddReq
	rpcReq.Productid = req.Productid
	rpcReq.Uid =uid

	_,err = l.svcCtx.UserRPC.AddUserCollection(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to UserCollectionAdd err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("error! UserCollectionAdd fail"),"rpc:Failed to UserCollectionAdd")
	}
	return &types.UserCollectionAddRes{},nil
}
