package mall

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hnchain/api/internal/logic/mall"
	"hnchain/api/internal/svc"
)

func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := mall.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
