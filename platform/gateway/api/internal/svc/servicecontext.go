package svc

import (
	"go-micro/platform/basic/rpc/basicclient"
	"go-micro/platform/gateway/api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	BasicRpc basicclient.Basic
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		BasicRpc: basicclient.NewBasic(zrpc.MustNewClient(c.BasicRpc)),
	}
}
