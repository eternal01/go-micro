package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemRegionModel = (*customSystemRegionModel)(nil)

type (
	// SystemRegionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemRegionModel.
	SystemRegionModel interface {
		systemRegionModel
	}

	customSystemRegionModel struct {
		*defaultSystemRegionModel
	}
)

// NewSystemRegionModel returns a model for the database table.
func NewSystemRegionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemRegionModel {
	return &customSystemRegionModel{
		defaultSystemRegionModel: newSystemRegionModel(conn, c, opts...),
	}
}
