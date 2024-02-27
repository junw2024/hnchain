package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRevAddrInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRevAddrInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRevAddrInfoLogic {
	return &GetUserRevAddrInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据主键id,查询收获地址
func (l *GetUserRevAddrInfoLogic) GetUserRevAddrInfo(in *user.UserRevAddrInfoReq) (*user.UserRevAddr, error) {
	revDrr, err := l.svcCtx.HnuserRevAddrModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("收获地址数据不存在"), "收获地址数据不存在")
		}
		return nil,err
	}

	var uRevAddr user.UserRevAddr
	copier.Copy(&uRevAddr,revDrr)

	return &uRevAddr, nil
}
