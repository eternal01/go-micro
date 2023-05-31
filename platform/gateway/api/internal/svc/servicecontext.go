package svc

import (
	"go-micro/platform/account/rpc/accountclient"
	"go-micro/platform/basic/rpc/basicclient"
	"go-micro/platform/file/rpc/fileclient"
	"go-micro/platform/gateway/api/internal/config"
	"go-micro/platform/information/rpc/informationclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	FileRpc        fileclient.File
	BasicRpc       basicclient.Basic
	AccountRpc     accountclient.Account
	InformationRpc informationclient.Information
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		FileRpc:        fileclient.NewFile(zrpc.MustNewClient(c.FileRpc)),
		BasicRpc:       basicclient.NewBasic(zrpc.MustNewClient(c.BasicRpc)),
		AccountRpc:     accountclient.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
		InformationRpc: informationclient.NewInformation(zrpc.MustNewClient(c.InformationRpc)),
	}
}
