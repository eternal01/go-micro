package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemRegionModel = (*customSystemRegionModel)(nil)

type (
	// SystemRegionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemRegionModel.
	SystemRegionModel interface {
		systemRegionModel

		RowBuilder() squirrel.SelectBuilder
		FindListBySelectBuilder(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemRegion, error)
		FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemRegion, error)
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

func (m *defaultSystemRegionModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(systemRegionRows).From(m.table)
}

func (m *defaultSystemRegionModel) FindListBySelectBuilder(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemRegion, error) {
	return []*SystemRegion{}, nil
}

func (m *defaultSystemRegionModel) FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemRegion, error) {
	var resp []*SystemRegion
	sql, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
