package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var (
	//文章未审核
	TopicStatusNotAudited int64 = 0
	//文章审核通过
	TopicStatusApproved int64 = 1
	//文章审核未通过
	TopicStatusRejection int64 = 2
)
