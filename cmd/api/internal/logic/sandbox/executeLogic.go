package sandbox

import (
	"context"
	"github.com/pkg/errors"
	"gyu-oj-sandbox/cmd/rpc/codesandbox"
	"gyu-oj-sandbox/common/xerr"

	"gyu-oj-sandbox/cmd/api/internal/svc"
	"gyu-oj-sandbox/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecuteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecuteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecuteLogic {
	return &ExecuteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecuteLogic) Execute(req *types.ExecuteReq) (resp *types.ExecuteResp, err error) {
	if len(req.InputList) <= 0 || req.Code == "" || req.Language == "" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.RequestParamError), "请求参数错误, inputList: %s, code: %s, language: %s", req.InputList, req.Code, req.Language)
	}

	sandboxResp, err := l.svcCtx.SandboxRpc.ExecuteCode(l.ctx, &codesandbox.ExecuteCodeReq{
		InputList: req.InputList,
		Code:      req.Code,
		Language:  req.Language,
	})
	if err != nil {
		return nil, err
	}

	return &types.ExecuteResp{
		OutputList:           sandboxResp.OutputList,
		Message:              sandboxResp.Message,
		Status:               sandboxResp.Status,
		ExecuteResultMessage: sandboxResp.ExecuteResultMessage,
		ExecuteResultTime:    sandboxResp.ExecuteResultTime,
		ExecuteResultMemory:  sandboxResp.ExecuteResultMemory,
	}, nil
}
