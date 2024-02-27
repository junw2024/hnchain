package mall

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hnchain/api/internal/logic/mall"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mall.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
