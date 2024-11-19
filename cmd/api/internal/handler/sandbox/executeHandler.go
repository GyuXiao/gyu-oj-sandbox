package sandbox

import (
	"gyu-oj-sandbox/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gyu-oj-sandbox/cmd/api/internal/logic/sandbox"
	"gyu-oj-sandbox/cmd/api/internal/svc"
	"gyu-oj-sandbox/cmd/api/internal/types"
)

// executeCode
func ExecuteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExecuteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sandbox.NewExecuteLogic(r.Context(), svcCtx)
		resp, err := l.Execute(&req)
		result.HttpResult(r, w, resp, err)
	}
}
