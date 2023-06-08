package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TopicModel = (*customTopicModel)(nil)

type (
	// TopicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicModel.
	TopicModel interface {
		topicModel

		CountBuilder() squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
		Count(ctx context.Context, rowBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Topic, error)
	}

	customTopicModel struct {
		*defaultTopicModel
	}
)

// NewTopicModel returns a model for the database table.
func NewTopicModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TopicModel {
	return &customTopicModel{
		defaultTopicModel: newTopicModel(conn, c, opts...),
	}
}

func (m *defaultTopicModel) CountBuilder() squirrel.SelectBuilder {
	return squirrel.Select("count(*)").From(m.table)
}

func (m *defaultTopicModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(topicRows).From(m.table)
}

func (m *defaultTopicModel) Count(ctx context.Context, rowBuilder squirrel.SelectBuilder) (int64, error) {
	var count int64
	sql, values, err := rowBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	err = m.QueryRowNoCacheCtx(ctx, &count, sql, values...)

	switch err {
	case nil:
		return count, nil
	default:
		return 0, err
	}
}

func (m *defaultTopicModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Topic, error) {
	var resp []*Topic
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
