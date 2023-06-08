package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuoteModel = (*customQuoteModel)(nil)

type (
	// QuoteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuoteModel.
	QuoteModel interface {
		quoteModel
	}

	customQuoteModel struct {
		*defaultQuoteModel
	}
)

// NewQuoteModel returns a model for the database table.
func NewQuoteModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuoteModel {
	return &customQuoteModel{
		defaultQuoteModel: newQuoteModel(conn, c, opts...),
	}
}
