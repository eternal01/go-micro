package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemIndustryModel = (*customSystemIndustryModel)(nil)

type (
	// SystemIndustryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemIndustryModel.
	SystemIndustryModel interface {
		systemIndustryModel

		FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemIndustry, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customSystemIndustryModel struct {
		*defaultSystemIndustryModel
	}
)

// NewSystemIndustryModel returns a model for the database table.
func NewSystemIndustryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemIndustryModel {
	return &customSystemIndustryModel{
		defaultSystemIndustryModel: newSystemIndustryModel(conn, c, opts...),
	}
}

func (m *defaultSystemIndustryModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(systemIndustryRows).From(m.table)
}

func (m *defaultSystemIndustryModel) FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemIndustry, error) {
	var resp []*SystemIndustry
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
