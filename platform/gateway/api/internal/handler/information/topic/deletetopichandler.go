package topic

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-micro/platform/gateway/api/internal/logic/information/topic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
)

func DeleteTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GatewayDeleteTopicRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := topic.NewDeleteTopicLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
