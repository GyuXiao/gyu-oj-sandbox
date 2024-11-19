// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	sandbox "gyu-oj-sandbox/cmd/api/internal/handler/sandbox"
	"gyu-oj-sandbox/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// executeCode
				Method:  http.MethodPost,
				Path:    "/sandbox/execute",
				Handler: sandbox.ExecuteHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gyu_oj/v1"),
	)
}
