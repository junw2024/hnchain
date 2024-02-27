package logic

import (
	"context"
	"time"

	"hnchain/common/xerr"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserRevAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserRevAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserRevAddrLogic {
	return &AddUserRevAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加收获地址
func (l *AddUserRevAddrLogic) AddUserRevAddr(in *user.UserRevAddrAddReq) (*user.UserRevAddrAddRes, error) {
	var addr  model.HnuserRevAddr
    _ = copier.Copy(&addr,in)
	addr.Createtime = time.Now()
	addr.Updatetime = time.Now()
	_,err := l.svcCtx.HnuserRevAddrModel.Insert(l.ctx,&addr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError),
		 "AddUserReceiveAddress Database Exception : %+v , err: %v", addr, err)
	}
	return &user.UserRevAddrAddRes{}, nil
}
