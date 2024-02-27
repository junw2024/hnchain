package logic

import (
	"context"
	"time"

	"hnchain/common/xerr"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserRevAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserRevAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserRevAddrLogic {
	return &DelUserRevAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收获地址
func (l *DelUserRevAddrLogic) DelUserRevAddr(in *user.UserRevAddrDelReq) (*user.UserCollectionDelRes, error) {
	_, err := l.svcCtx.HnuserRevAddrModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "没有该收获地址")
		}
		return nil, err
	}

	addr := new(model.HnuserRevAddr)
	addr.Id = in.Id
	addr.Isdelete = true
	addr.Updatetime = time.Now()

	err = l.svcCtx.HnuserRevAddrModel.UpdateIsDelete(l.ctx, addr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DelUserRevAddress Database Exception : %+v , err: %v", addr, err)
	}
	return &user.UserCollectionDelRes{}, nil
}
