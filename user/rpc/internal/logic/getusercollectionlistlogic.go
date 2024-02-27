package logic

import (
	"context"

	"hnchain/common/xerr"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏列表
func (l *GetUserCollectionListLogic) GetUserCollectionList(in *user.UserCollectionListReq) (*user.UserCollectionListRes, error) {
	collectionList, err := l.svcCtx.HnuserCollectionModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError),
			"Failed  get user's Collection list err : %v , in :%+v", err, in)
	}

	var resp []int64
	for _, it := range collectionList {
		resp = append(resp, it.Productid)
	}
	return &user.UserCollectionListRes{List: resp}, nil
}
