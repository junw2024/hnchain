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
	"google.golang.org/grpc/status"
)

type EditUserRevAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserRevAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserRevAddrLogic {
	return &EditUserRevAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 编辑收获地址
func (l *EditUserRevAddrLogic) EditUserRevAddr(in *user.UserRevAddrEditReq) (*user.UserRevAddrEditRes, error) {
	oAddr, err := l.svcCtx.HnuserRevAddrModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "数据不存在")
		}
		return nil, err
	}
	var addr model.HnuserRevAddr
	errcopy := copier.Copy(&addr,in)
	if errcopy != nil {
		return nil, errcopy
	}
	
	addr.Uid = oAddr.Uid
	addr.Updatetime = time.Now()
	addr.Createtime = oAddr.Createtime

	err = l.svcCtx.HnuserRevAddrModel.EditUserRevAddr(l.ctx,&addr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "EditUserRevAddr Database Exception : %+v , err: %v", addr, err)
	}
	return &user.UserRevAddrEditRes{}, nil
}
