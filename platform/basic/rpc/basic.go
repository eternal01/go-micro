package main

import (
	"flag"
	"fmt"

	"go-micro/common/interceptor/rpcserver"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/config"
	basiccategoryserviceServer "go-micro/platform/basic/rpc/internal/server/basiccategoryservice"
	basicclassifyserviceServer "go-micro/platform/basic/rpc/internal/server/basicclassifyservice"
	basicconfigurationserviceServer "go-micro/platform/basic/rpc/internal/server/basicconfigurationservice"
	basicindustryserviceServer "go-micro/platform/basic/rpc/internal/server/basicindustryservice"
	basicregionserviceServer "go-micro/platform/basic/rpc/internal/server/basicregionservice"
	basicserviceServer "go-micro/platform/basic/rpc/internal/server/basicservice"
	basicstageserviceServer "go-micro/platform/basic/rpc/internal/server/basicstageservice"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/basic.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		basic.RegisterBasicServiceServer(grpcServer, basicserviceServer.NewBasicServiceServer(ctx))
		basic.RegisterBasicRegionServiceServer(grpcServer, basicregionserviceServer.NewBasicRegionServiceServer(ctx))
		basic.RegisterBasicIndustryServiceServer(grpcServer, basicindustryserviceServer.NewBasicIndustryServiceServer(ctx))
		basic.RegisterBasicClassifyServiceServer(grpcServer, basicclassifyserviceServer.NewBasicClassifyServiceServer(ctx))
		basic.RegisterBasicCategoryServiceServer(grpcServer, basiccategoryserviceServer.NewBasicCategoryServiceServer(ctx))
		basic.RegisterBasicStageServiceServer(grpcServer, basicstageserviceServer.NewBasicStageServiceServer(ctx))
		basic.RegisterBasicConfigurationServiceServer(grpcServer, basicconfigurationserviceServer.NewBasicConfigurationServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
