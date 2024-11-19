package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gyu-oj-sandbox/cmd/api/internal/config"
	"gyu-oj-sandbox/cmd/rpc/codesandbox"
)

type ServiceContext struct {
	Config     config.Config
	SandboxRpc codesandbox.CodeSandbox
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		SandboxRpc: codesandbox.NewCodeSandbox(zrpc.MustNewClient(c.SandboxRpcConf)),
	}
}
