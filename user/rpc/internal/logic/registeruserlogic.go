package logic

import (
	"context"
	"time"

	"hnchain/common/tool"
	"hnchain/common/xerr"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册用户
func (l *RegisterUserLogic) RegisterUser(in *user.RegisterUserReq) (*user.RegisterUserRes, error) {
	_,err := l.svcCtx.HnuserModel.FindOneByUsername(l.ctx,in.Username)
	if err == nil {
		return nil,errors.Wrap(xerr.NewErrMsg("用户名已注册"), "注册用户已存在")
	}
	if err != nil && err != model.ErrNotFound {
		return nil,errors.Wrap(xerr.NewErrMsg("注册异常"), "数据库异常了")
	}

	_, err = l.svcCtx.HnuserModel.FindOneByPhone(l.ctx,in.Phone)
	if err == nil {
		return nil,errors.Wrap(xerr.NewErrMsg("手机号已注册"), "手机号已存在")
	}
	if err != nil && err != model.ErrNotFound {
		return nil,errors.Wrap(xerr.NewErrMsg("注册异常"), "数据库异常了")
	}
    
	//注册用户
    var nUser model.Hnuser
	nUser.Username = in.Username
	nUser.Phone = in.Phone
	nUser.Password,_ = tool.Md5ByString(in.Password)
	nUser.Sex="0"
	nUser.Updatetime =time.Now()
	nUser.Createtime = time.Now()
	_,err = l.svcCtx.HnuserModel.Insert(l.ctx,&nUser)
	if err != nil {
		return nil,errors.Wrap(xerr.NewErrMsg("注册异常"), "数据库异常了")
	}
	return &user.RegisterUserRes{}, nil
}
