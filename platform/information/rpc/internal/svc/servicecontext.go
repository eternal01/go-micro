package svc

import (
	"go-micro/platform/information/model"
	"go-micro/platform/information/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	TopicModel            model.TopicModel
	TopicAuditRecordModel model.TopicAuditRecordModel
	NewsModel             model.NewsModel
	ExhibitionModel       model.ExhibitionModel
	MiningRightModel      model.MiningRightModel
	QuoteModel            model.QuoteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
		TopicModel:            model.NewTopicModel(conn, c.CacheRedis),
		TopicAuditRecordModel: model.NewTopicAuditRecordModel(conn, c.CacheRedis),
		NewsModel:             model.NewNewsModel(conn, c.CacheRedis),
		ExhibitionModel:       model.NewExhibitionModel(conn, c.CacheRedis),
		MiningRightModel:      model.NewMiningRightModel(conn, c.CacheRedis),
		QuoteModel:            model.NewQuoteModel(conn, c.CacheRedis),
	}
}
