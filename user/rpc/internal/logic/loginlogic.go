package logic

import (
	"context"
	"hnchain/common/tool"
	"hnchain/common/xerr"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/model"
	"hnchain/user/rpc/user"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/jinzhu/copier"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	//verify user exists
    userDb,err := l.svcCtx.HnuserModel.FindOneByUsername(l.ctx,in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil,errors.Wrapf(xerr.NewErrCode(xerr.DbError),
			"根据username查询用户信息失败，username:%s,err:%v", in.Username, err)
		}
		return nil,err
	}

	//verify user password
	md5Str,_ :=  tool.Md5ByString(in.Password)
	if !(md5Str == userDb.Password) {
		return nil, errors.Wrap(xerr.NewErrMsg("账号或密码错误"), "密码错误")
	}

	//return sql
	var resp  user.LoginResponse
	_ = copier.Copy(&resp,userDb)
	
	return &resp, nil
}
