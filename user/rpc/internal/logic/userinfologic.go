package logic

import (
	"context"

	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得用户信息
func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	uMember, err := l.svcCtx.HnuserModel.FindOne(l.ctx, in.Id)
	l.Logger.Infof("uMember:%v",uMember)

	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil,status.Error(500,err.Error())
	}
	var resp user.UserInfo
	_ = copier.Copy(&resp,uMember)
	resp.CreateTime = uMember.Createtime.Unix();
	resp.UpdateTime = uMember.Createtime.Unix();

	return &user.UserInfoResponse{
		User: &resp,
	}, nil
}
