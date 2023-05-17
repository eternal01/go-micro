package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersAuthsModel = (*customUsersAuthsModel)(nil)

type (
	// UsersAuthsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersAuthsModel.
	UsersAuthsModel interface {
		usersAuthsModel

		TxInsert(ctx context.Context, tx *sql.Tx, data *UsersAuths) (sql.Result, error)
	}

	customUsersAuthsModel struct {
		*defaultUsersAuthsModel
	}
)

// NewUsersAuthsModel returns a model for the database table.
func NewUsersAuthsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersAuthsModel {
	return &customUsersAuthsModel{
		defaultUsersAuthsModel: newUsersAuthsModel(conn, c, opts...),
	}
}

func (m *defaultUsersAuthsModel) TxInsert(ctx context.Context, tx *sql.Tx, data *UsersAuths) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, usersAuthsRowsExpectAutoSet)
	return tx.ExecContext(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.AuthCredential, data.DeletedAt)
}
