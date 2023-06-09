package response

import (
	"go-micro/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError:
			body.Code = e.Code
			body.Msg = e.Msg
		default:
			body.Code = -1
			body.Msg = err.Error()
		}
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	logx.WithContext(r.Context()).Slowf("Response - %+v", resp)
	logx.CollectSysLog()
	httpx.OkJson(w, body)
}
