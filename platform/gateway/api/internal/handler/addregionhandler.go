package handler

import (
	"net/http"

	"go-micro/common/response"
	"go-micro/platform/gateway/api/internal/logic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func addRegionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GatewayAddRegionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddRegionLogic(r.Context(), svcCtx)
		resp, err := l.AddRegion(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r, w, resp, err)
	}
}
