package mall

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hnchain/api/internal/logic/mall"
	"hnchain/api/internal/svc"
)

func HomeBannerHanderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := mall.NewHomeBannerHanderLogic(r.Context(), svcCtx)
		resp, err := l.HomeBannerHander()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
