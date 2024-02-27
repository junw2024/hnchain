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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) (*types.UserRegisterRes,  error) {
	var rpcReq userclient.RegisterUserReq
	_ = copier.Copy(&rpcReq,req)
	_,err := l.svcCtx.UserRPC.RegisterUser(l.ctx,&rpcReq)
	if err != nil {
		l.Logger.Errorf("rpc:Failed to Register: %v ,req:%+v", err, req)
		return nil, errors.Wrapf(xerr.NewErrMsg("rpc:Failed to Register!"), "req: %+v", req)
	}
	return &types.UserRegisterRes{},nil
}
