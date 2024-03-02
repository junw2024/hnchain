package mall

import (
	"hnchain/api/internal/logic/mall"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/result"
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := mall.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
