package result

import (
	"fmt"
	"hnchain/common/xerr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// http response
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errcode := xerr.ServerCommonError
		//default error msg
		errmsg := "服务器开小差啦，稍后再试"

		causeErr := errors.Cause(err)

		if e, ok := causeErr.(*xerr.CodeError); ok {
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {

				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) {
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	}
}

// http auth error
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errcode := xerr.ServerCommonError
		errmsg := "服务器开小差啦，稍后再来试一试"
		
		causeErr := errors.Cause(err)

		if e, ok := causeErr.(*xerr.CodeError); ok {
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			//是否grpc 错误
			if gstatus, ok := status.FromError(causeErr); ok {
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) {
					errcode = grpcCode
					errmsg = gstatus.Message()
				}

			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

// http param error（参数错误）
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.ReuqestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.ReuqestParamError, errMsg))
}
