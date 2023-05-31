package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	FileRpc        zrpc.RpcClientConf
	BasicRpc       zrpc.RpcClientConf
	AccountRpc     zrpc.RpcClientConf
	InformationRpc zrpc.RpcClientConf

	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
