package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersAuthsModel = (*customUsersAuthsModel)(nil)

type (
	// UsersAuthsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersAuthsModel.
	UsersAuthsModel interface {
		usersAuthsModel
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
