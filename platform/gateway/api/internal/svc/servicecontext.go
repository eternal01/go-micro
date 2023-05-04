package svc

import (
	"go-micro/platform/account/rpc/accountclient"
	"go-micro/platform/basic/rpc/basicclient"
	"go-micro/platform/gateway/api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	BasicRpc   basicclient.Basic
	AccountRpc accountclient.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		BasicRpc:   basicclient.NewBasic(zrpc.MustNewClient(c.BasicRpc)),
		AccountRpc: accountclient.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
	}
}
