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

type DelUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserCollectionLogic {
	return &DelUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收藏
func (l *DelUserCollectionLogic) DelUserCollection(in *user.UserCollectionDelReq) (*user.UserCollectionDelRes, error) {
	//查询
	_, err := l.svcCtx.HnuserCollectionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "该商品没有被收藏")
		}
		return nil,err
	}
	
	collection := new(model.HnuserCollection)
	collection.Id = in.Id
	collection.Isdelete = true
	collection.Updatetime = time.Now()
	l.svcCtx.HnuserCollectionModel.UpdateIsDelete(l.ctx, collection)
	return &user.UserCollectionDelRes{}, nil
}
