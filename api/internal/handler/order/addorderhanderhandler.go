package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hnchain/api/internal/logic/order"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
)

func AddOrderHanderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewAddOrderHanderLogic(r.Context(), svcCtx)
		resp, err := l.AddOrderHander(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
