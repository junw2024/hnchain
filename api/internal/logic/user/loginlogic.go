package user

import (
	"context"
	"time"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/jwtx"
	"hnchain/common/xerr"
	"hnchain/user/rpc/userclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	var loginReq userclient.LoginRequest
	loginReq.Username = req.Username
	loginReq.Password = req.Password

	l.Logger.Info("LoginLogic..................")

	res, err := l.svcCtx.UserRPC.Login(l.ctx, &loginReq)

	if err != nil {
		l.Logger.Errorf("rpc:Failed to Login err : %v ,req:%+v", err, req)
		return nil, errors.Wrapf(xerr.NewErrMsg("Login error!"), "req: %+v", req)
	}

	//generate token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	
	accessToken, err := jwtx.GetToken(accessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil,err
	}
	
	return &types.LoginRes{
		AccessToken: accessToken,
		AccessExpire: now+accessExpire,
	},nil
}
