package mall

import (
	"net/http"

	"hnchain/api/internal/logic/mall"
	"hnchain/api/internal/svc"
	"hnchain/api/internal/types"
	"hnchain/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r,w,err)
			return
		}

		l := mall.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		
		result.HttpResult(r,w,resp,err)
	}
}
