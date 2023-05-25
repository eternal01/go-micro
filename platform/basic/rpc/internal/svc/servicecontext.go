package svc

import (
	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RegionModel   model.SystemRegionModel
	IndustryModel model.SystemIndustryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		RegionModel:   model.NewSystemRegionModel(conn, c.CacheRedis),
		IndustryModel: model.NewSystemIndustryModel(conn, c.CacheRedis),
	}
}
