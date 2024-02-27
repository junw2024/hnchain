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

type GetUserRevAddrListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRevAddrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRevAddrListLogic {
	return &GetUserRevAddrListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取收获地址列表
func (l *GetUserRevAddrListLogic) GetUserRevAddrList(in *user.UserRevAddrListReq) (*user.UserRevAddrListRes, error) {

	revList, err := l.svcCtx.HnuserRevAddrModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed  get user's receive address list err :%v,in:%+v",err,in)
	}

	var resp []*user.UserRevAddr
	for _,it := range revList {
		var addr user.UserRevAddr
		_ = copier.Copy(&addr,it)
		resp =append(resp, &addr)
	}
	return &user.UserRevAddrListRes{List: resp}, nil
}
