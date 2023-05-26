package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemCategoryModel = (*customSystemCategoryModel)(nil)

type (
	// SystemCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemCategoryModel.
	SystemCategoryModel interface {
		systemCategoryModel

		RowBuilder() squirrel.SelectBuilder
		FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemCategory, error)
	}

	customSystemCategoryModel struct {
		*defaultSystemCategoryModel
	}
)

// NewSystemCategoryModel returns a model for the database table.
func NewSystemCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemCategoryModel {
	return &customSystemCategoryModel{
		defaultSystemCategoryModel: newSystemCategoryModel(conn, c, opts...),
	}
}

func (m *defaultSystemCategoryModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(systemCategoryRows).From(m.table)
}

func (m *defaultSystemCategoryModel) FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemCategory, error) {
	var resp []*SystemCategory
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
