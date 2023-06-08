package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsModel = (*customNewsModel)(nil)

type (
	// NewsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsModel.
	NewsModel interface {
		newsModel
	}

	customNewsModel struct {
		*defaultNewsModel
	}
)

// NewNewsModel returns a model for the database table.
func NewNewsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NewsModel {
	return &customNewsModel{
		defaultNewsModel: newNewsModel(conn, c, opts...),
	}
}
