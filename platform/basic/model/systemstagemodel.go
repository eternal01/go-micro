package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemStageModel = (*customSystemStageModel)(nil)

type (
	// SystemStageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemStageModel.
	SystemStageModel interface {
		systemStageModel

		RowBuilder() squirrel.SelectBuilder
		FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemStage, error)
	}

	customSystemStageModel struct {
		*defaultSystemStageModel
	}
)

// NewSystemStageModel returns a model for the database table.
func NewSystemStageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemStageModel {
	return &customSystemStageModel{
		defaultSystemStageModel: newSystemStageModel(conn, c, opts...),
	}
}

func (m *defaultSystemStageModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(systemClassifyRows).From(m.table)
}

func (m *defaultSystemStageModel) FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemStage, error) {
	var resp []*SystemStage
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
