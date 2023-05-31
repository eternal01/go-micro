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
	ClassifyModel model.SystemClassifyModel
	CategoryModel model.SystemCategoryModel
	StageModel    model.SystemStageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		RegionModel:   model.NewSystemRegionModel(conn, c.CacheRedis),
		IndustryModel: model.NewSystemIndustryModel(conn, c.CacheRedis),
		ClassifyModel: model.NewSystemClassifyModel(conn, c.CacheRedis),
		CategoryModel: model.NewSystemCategoryModel(conn, c.CacheRedis),
		StageModel:    model.NewSystemStageModel(conn, c.CacheRedis),
	}
}
