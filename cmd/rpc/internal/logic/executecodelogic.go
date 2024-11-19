package logic

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/zeromicro/go-zero/core/logc"

	"gyu-oj-sandbox/cmd/rpc/internal/svc"
	"gyu-oj-sandbox/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecuteCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExecuteCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecuteCodeLogic {
	return &ExecuteCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExecuteCodeLogic) ExecuteCode(in *pb.ExecuteCodeReq) (*pb.ExecuteCodeResp, error) {
	var sandbox ExecuteCodeItf
	// 1,new 一个代码沙箱
	switch l.svcCtx.Config.SandboxBy.Type {
	case "golang":
		sandbox = NewSandboxByGoNative(l.ctx)
	case "docker":
		sandbox = NewSandboxByDocker(l.ctx, l.svcCtx.DockerClient)
	}

	// 2,使用代码沙箱
	resp, err := SandboxTemplate(sandbox, in)
	// 开启 goroutine 释放资源
	go ReleaseSource(context.Background(), l.svcCtx.Config.SandboxBy.Type, l.svcCtx.DockerClient)

	if err != nil {
		return nil, err
	}

	// 3,返回代码输出结果
	return resp, nil
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
