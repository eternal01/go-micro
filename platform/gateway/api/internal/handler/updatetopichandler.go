package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-micro/platform/gateway/api/internal/logic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
)

func updateTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GatewayUpdateTopicRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateTopicLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
