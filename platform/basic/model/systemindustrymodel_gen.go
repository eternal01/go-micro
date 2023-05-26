// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	systemIndustryFieldNames          = builder.RawFieldNames(&SystemIndustry{})
	systemIndustryRows                = strings.Join(systemIndustryFieldNames, ",")
	systemIndustryRowsExpectAutoSet   = strings.Join(stringx.Remove(systemIndustryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	systemIndustryRowsWithPlaceHolder = strings.Join(stringx.Remove(systemIndustryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSystemIndustryIdPrefix         = "cache:systemIndustry:id:"
	cacheSystemIndustryIndustryIdPrefix = "cache:systemIndustry:industryId:"
)

type (
	systemIndustryModel interface {
		Insert(ctx context.Context, data *SystemIndustry) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SystemIndustry, error)
		FindOneByIndustryId(ctx context.Context, industryId string) (*SystemIndustry, error)
		Update(ctx context.Context, data *SystemIndustry) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSystemIndustryModel struct {
		sqlc.CachedConn
		table string
	}

	SystemIndustry struct {
		Id          int64  `db:"id"`          // 主键id
		IndustryId  string `db:"industry_id"` // 行业id
		Name        string `db:"name"`        // 行业名称
		ParentId    string `db:"parent_id"`   // 父级id
		LevelType   int64  `db:"level_type"`  // 等级
		Description string `db:"description"` // 描述
	}
)

func newSystemIndustryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSystemIndustryModel {
	return &defaultSystemIndustryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`system_industry`",
	}
}

func (m *defaultSystemIndustryModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	systemIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIdPrefix, id)
	systemIndustryIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIndustryIdPrefix, data.IndustryId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, systemIndustryIdKey, systemIndustryIndustryIdKey)
	return err
}

func (m *defaultSystemIndustryModel) FindOne(ctx context.Context, id int64) (*SystemIndustry, error) {
	systemIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIdPrefix, id)
	var resp SystemIndustry
	err := m.QueryRowCtx(ctx, &resp, systemIndustryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", systemIndustryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemIndustryModel) FindOneByIndustryId(ctx context.Context, industryId string) (*SystemIndustry, error) {
	systemIndustryIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIndustryIdPrefix, industryId)
	var resp SystemIndustry
	err := m.QueryRowIndexCtx(ctx, &resp, systemIndustryIndustryIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `industry_id` = ? limit 1", systemIndustryRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, industryId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemIndustryModel) Insert(ctx context.Context, data *SystemIndustry) (sql.Result, error) {
	systemIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIdPrefix, data.Id)
	systemIndustryIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIndustryIdPrefix, data.IndustryId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, systemIndustryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.IndustryId, data.Name, data.ParentId, data.LevelType, data.Description)
	}, systemIndustryIdKey, systemIndustryIndustryIdKey)
	return ret, err
}

func (m *defaultSystemIndustryModel) Update(ctx context.Context, newData *SystemIndustry) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	systemIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIdPrefix, data.Id)
	systemIndustryIndustryIdKey := fmt.Sprintf("%s%v", cacheSystemIndustryIndustryIdPrefix, data.IndustryId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, systemIndustryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.IndustryId, newData.Name, newData.ParentId, newData.LevelType, newData.Description, newData.Id)
	}, systemIndustryIdKey, systemIndustryIndustryIdKey)
	return err
}

func (m *defaultSystemIndustryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSystemIndustryIdPrefix, primary)
}

func (m *defaultSystemIndustryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", systemIndustryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSystemIndustryModel) tableName() string {
	return m.table
}