package svc

import (
	"go-micro/platform/account/rpc/accountclient"
	"go-micro/platform/basic/rpc/client/basiccategoryservice"
	"go-micro/platform/basic/rpc/client/basicclassifyservice"
	"go-micro/platform/basic/rpc/client/basicconfigurationservice"
	"go-micro/platform/basic/rpc/client/basicindustryservice"
	"go-micro/platform/basic/rpc/client/basicregionservice"
	"go-micro/platform/basic/rpc/client/basicservice"
	"go-micro/platform/basic/rpc/client/basicstageservice"
	"go-micro/platform/file/rpc/fileclient"
	"go-micro/platform/gateway/api/internal/config"
	"go-micro/platform/information/rpc/client/informationservice"
	"go-micro/platform/information/rpc/client/informationtopicauditrecordservice"
	"go-micro/platform/information/rpc/client/informationtopicservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc fileclient.File

	BasicCategoryRpc      basiccategoryservice.BasicCategoryService
	BasicClassifyRpc      basicclassifyservice.BasicClassifyService
	BasicConfigurationRpc basicconfigurationservice.BasicConfigurationService
	BasicIndustryRpc      basicindustryservice.BasicIndustryService
	BasicRegionRpc        basicregionservice.BasicRegionService
	BasicRpc              basicservice.BasicService
	BasicStageRpc         basicstageservice.BasicStageService

	AccountRpc accountclient.Account

	InformationRpc                 informationservice.InformationService
	InformationTopicRpc            informationtopicservice.InformationTopicService
	InformationTopicAuditRecordRpc informationtopicauditrecordservice.InformationTopicAuditRecordService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FileRpc: fileclient.NewFile(zrpc.MustNewClient(c.FileRpc)),

		BasicCategoryRpc:      basiccategoryservice.NewBasicCategoryService(zrpc.MustNewClient(c.BasicRpc)),
		BasicClassifyRpc:      basicclassifyservice.NewBasicClassifyService(zrpc.MustNewClient(c.BasicRpc)),
		BasicConfigurationRpc: basicconfigurationservice.NewBasicConfigurationService(zrpc.MustNewClient(c.BasicRpc)),
		BasicIndustryRpc:      basicindustryservice.NewBasicIndustryService(zrpc.MustNewClient(c.BasicRpc)),
		BasicRegionRpc:        basicregionservice.NewBasicRegionService(zrpc.MustNewClient(c.BasicRpc)),
		BasicRpc:              basicservice.NewBasicService(zrpc.MustNewClient(c.BasicRpc)),
		BasicStageRpc:         basicstageservice.NewBasicStageService(zrpc.MustNewClient(c.BasicRpc)),

		AccountRpc: accountclient.NewAccount(zrpc.MustNewClient(c.AccountRpc)),

		InformationRpc:                 informationservice.NewInformationService(zrpc.MustNewClient(c.InformationRpc)),
		InformationTopicRpc:            informationtopicservice.NewInformationTopicService(zrpc.MustNewClient(c.InformationRpc)),
		InformationTopicAuditRecordRpc: informationtopicauditrecordservice.NewInformationTopicAuditRecordService(zrpc.MustNewClient(c.InformationRpc)),
	}
}
