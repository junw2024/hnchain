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

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoRes, err error) {
	
	uid,_ := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		l.Logger.Errorf("token:Failed to AddReveAddr err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("Detail error! Get token"),"")
	}

	var rpcReq userclient.UserInfoRequest
	rpcReq.Id=uid
	userInfo, err := l.svcCtx.UserRPC.UserInfo(l.ctx,&rpcReq)
	
	if err != nil {
		l.Logger.Errorf("rpc:Failed to Detail err : %v ,req:%+v", err, req)
		return nil,errors.Wrap(xerr.NewErrMsg("Failed to userDetail err"),"rpc userDetail err")
	}

	var user types.UserInfo
	_ = copier.Copy(&user,userInfo.User)

	return &types.UserInfoRes{UserInfo: user},nil
}
