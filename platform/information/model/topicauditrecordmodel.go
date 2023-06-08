package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TopicAuditRecordModel = (*customTopicAuditRecordModel)(nil)

type (
	// TopicAuditRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicAuditRecordModel.
	TopicAuditRecordModel interface {
		topicAuditRecordModel

		RowBuilder() squirrel.SelectBuilder
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*TopicAuditRecord, error)
	}

	customTopicAuditRecordModel struct {
		*defaultTopicAuditRecordModel
	}
)

// NewTopicAuditRecordModel returns a model for the database table.
func NewTopicAuditRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TopicAuditRecordModel {
	return &customTopicAuditRecordModel{
		defaultTopicAuditRecordModel: newTopicAuditRecordModel(conn, c, opts...),
	}
}

func (m *defaultTopicAuditRecordModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(topicAuditRecordRows).From(m.table)
}

func (m *defaultTopicAuditRecordModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*TopicAuditRecord, error) {
	var resp []*TopicAuditRecord
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
