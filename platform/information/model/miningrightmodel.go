package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MiningRightModel = (*customMiningRightModel)(nil)

type (
	// MiningRightModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMiningRightModel.
	MiningRightModel interface {
		miningRightModel
	}

	customMiningRightModel struct {
		*defaultMiningRightModel
	}
)

// NewMiningRightModel returns a model for the database table.
func NewMiningRightModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MiningRightModel {
	return &customMiningRightModel{
		defaultMiningRightModel: newMiningRightModel(conn, c, opts...),
	}
}
