package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// rest.RestConf
	// UserRpc zrpc.RpcClientConf
}
