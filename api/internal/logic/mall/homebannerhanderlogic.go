package mall

import (
	"context"

	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBannerHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBannerHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerHanderLogic {
	return &HomeBannerHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBannerHanderLogic) HomeBannerHander() (resp *types.HomeBannerRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
