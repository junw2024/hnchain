package logic

import (
	"context"

	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginTstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginTstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginTstLogic {
	return &LoginTstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginTstLogic) LoginTst(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{}, nil
}
