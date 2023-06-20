package main

import (
	"flag"
	"fmt"

	"go-micro/common/interceptor/rpcserver"
	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/config"
	informationserviceServer "go-micro/platform/information/rpc/internal/server/informationservice"
	informationtopicauditrecordserviceServer "go-micro/platform/information/rpc/internal/server/informationtopicauditrecordservice"
	informationtopicserviceServer "go-micro/platform/information/rpc/internal/server/informationtopicservice"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/information.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		information.RegisterInformationServiceServer(grpcServer, informationserviceServer.NewInformationServiceServer(ctx))
		information.RegisterInformationTopicServiceServer(grpcServer, informationtopicserviceServer.NewInformationTopicServiceServer(ctx))
		information.RegisterInformationTopicAuditRecordServiceServer(grpcServer, informationtopicauditrecordserviceServer.NewInformationTopicAuditRecordServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
