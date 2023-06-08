package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExhibitionModel = (*customExhibitionModel)(nil)

type (
	// ExhibitionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExhibitionModel.
	ExhibitionModel interface {
		exhibitionModel
	}

	customExhibitionModel struct {
		*defaultExhibitionModel
	}
)

// NewExhibitionModel returns a model for the database table.
func NewExhibitionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExhibitionModel {
	return &customExhibitionModel{
		defaultExhibitionModel: newExhibitionModel(conn, c, opts...),
	}
}
