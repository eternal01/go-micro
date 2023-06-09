// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-micro/platform/gateway/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/region/:id",
				Handler: getRegionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/regions/:parent_id",
				Handler: getRegionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/region",
				Handler: addRegionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: registerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/gateway"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/:id",
				Handler: getUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/:id",
				Handler: updateUserHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/gateway"),
	)
}
