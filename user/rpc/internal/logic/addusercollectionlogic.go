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

type AddUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserCollectionLogic {
	return &AddUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加收藏
func (l *AddUserCollectionLogic) AddUserCollection(in *user.UserCollectionAddReq) (*user.UserCollectionAddRes, error) {
	collection := new (model.HnuserCollection)
	collection.Uid = in.Uid
	collection.Productid = in.Productid
    collection.Createtime = time.Now()
	collection.Updatetime =  time.Now()
	_,err := l.svcCtx.HnuserCollectionModel.Insert(l.ctx,collection)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUserCollection Database Exception : %+v , err: %v", collection, err)
	}
	return &user.UserCollectionAddRes{}, nil
}
