package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemConfigurationModel = (*customSystemConfigurationModel)(nil)

type (
	// SystemConfigurationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemConfigurationModel.
	SystemConfigurationModel interface {
		systemConfigurationModel

		RowBuilder() squirrel.SelectBuilder
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemConfiguration, error)
	}

	customSystemConfigurationModel struct {
		*defaultSystemConfigurationModel
	}
)

// NewSystemConfigurationModel returns a model for the database table.
func NewSystemConfigurationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemConfigurationModel {
	return &customSystemConfigurationModel{
		defaultSystemConfigurationModel: newSystemConfigurationModel(conn, c, opts...),
	}
}

func (m *defaultSystemConfigurationModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(systemConfigurationRows).From(m.table)
}

func (m *defaultSystemConfigurationModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemConfiguration, error) {
	var resp []*SystemConfiguration
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
