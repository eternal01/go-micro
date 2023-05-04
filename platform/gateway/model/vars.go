package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const Enable int64 = 1
const Disable int64 = 2

const AuthTypeSystem int64 = 0
