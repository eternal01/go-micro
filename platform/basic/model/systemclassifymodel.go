package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemClassifyModel = (*customSystemClassifyModel)(nil)

type (
	// SystemClassifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemClassifyModel.
	SystemClassifyModel interface {
		systemClassifyModel

		RowBuilder() squirrel.SelectBuilder
		FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemClassify, error)
	}

	customSystemClassifyModel struct {
		*defaultSystemClassifyModel
	}
)

// NewSystemClassifyModel returns a model for the database table.
func NewSystemClassifyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemClassifyModel {
	return &customSystemClassifyModel{
		defaultSystemClassifyModel: newSystemClassifyModel(conn, c, opts...),
	}
}

func (m *defaultSystemClassifyModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(systemClassifyRows).From(m.table)
}

func (m *defaultSystemClassifyModel) FindChildrenList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SystemClassify, error) {
	var resp []*SystemClassify
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
