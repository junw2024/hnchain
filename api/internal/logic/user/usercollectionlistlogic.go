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

type UserCollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionListLogic {
	return &UserCollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionListLogic) UserCollectionList(req *types.UserCollectionListReq) (resp *types.UserCollectionListRes, err error) {
	uid,_ := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		l.Logger.Errorf("token:Failed to UserCollectionList err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("UserCollectionList error! Get token"),"")
	}

	var rpcReq userclient.UserCollectionListReq
	rpcReq.Uid=uid

    rpcRes,err := l.svcCtx.UserRPC.GetUserCollectionList(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to UserCollectionList err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("UserCollectionList error! Get List"),"")
	}

	var resList []int64
	resList = append(resList, rpcRes.List...)
	return &types.UserCollectionListRes{Productid: resList},nil
}
