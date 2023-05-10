package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel

		TxInsert(ctx context.Context, tx *sql.Tx, data *Users) (sql.Result, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

func (m *defaultUsersModel) TxInsert(ctx context.Context, tx *sql.Tx, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	return tx.ExecContext(ctx, query, data.UserName, data.NickName, data.Avatar, data.Gender, data.Mobile, data.Email, data.Status, data.DeletedAt)
}
