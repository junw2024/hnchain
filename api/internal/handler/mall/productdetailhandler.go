package mall

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hnchain/api/internal/logic/mall"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
)

func ProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mall.NewProductDetailLogic(r.Context(), svcCtx)
		resp, err := l.ProductDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
