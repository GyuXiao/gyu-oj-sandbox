package sandbox

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
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

	var sandbox ExecuteCodeItf
	// 1,new 一个代码沙箱
	switch l.svcCtx.Config.SandboxBy.Type {
	case "golang":
		sandbox = NewSandboxByGoNative(l.ctx)
	case "docker":
		sandbox = NewSandboxByDocker(l.ctx, l.svcCtx.DockerClient)
	}

	// 2,使用代码沙箱
	resp, err = SandboxTemplate(sandbox, req)
	// 开启 goroutine 释放资源
	go ReleaseSource(context.Background(), l.svcCtx.Config.SandboxBy.Type, l.svcCtx.DockerClient)

	if err != nil {
		return nil, err
	}

	return &types.ExecuteResp{
		OutputList:           resp.OutputList,
		Message:              resp.Message,
		Status:               resp.Status,
		ExecuteResultMessage: resp.ExecuteResultMessage,
		ExecuteResultTime:    resp.ExecuteResultTime,
		ExecuteResultMemory:  resp.ExecuteResultMemory,
	}, nil
}

func ReleaseSource(ctx context.Context, sandboxType string, cli *client.Client) {
	if sandboxType != "docker" || GlobalContainerID == "" || cli == nil {
		return
	}
	// 删除容器（先停后删）
	err := cli.ContainerStop(ctx, GlobalContainerID, container.StopOptions{})
	if err != nil {
		logc.Infof(ctx, "停止容器错误: %v", err)
	}
	err = cli.ContainerRemove(ctx, GlobalContainerID, container.RemoveOptions{})
	if err != nil {
		logc.Infof(ctx, "删除容器错误: %v", err)
	}
}
