package user

import (
	"net/http"

	"hnchain/api/internal/logic/user"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelRevAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRevAddrDelReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r,w,err)
			return
		}

		l := user.NewDelRevAddressLogic(r.Context(), svcCtx)
		resp, err := l.DelRevAddress(&req)
		result.HttpResult(r,w,resp,err)
	}
}
