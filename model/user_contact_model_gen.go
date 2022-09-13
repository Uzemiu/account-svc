// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userContactFieldNames          = builder.RawFieldNames(&UserContact{})
	userContactRows                = strings.Join(userContactFieldNames, ",")
	userContactRowsExpectAutoSet   = strings.Join(stringx.Remove(userContactFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userContactRowsWithPlaceHolder = strings.Join(stringx.Remove(userContactFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	userContactModel interface {
		Insert(ctx context.Context, data *UserContact) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserContact, error)
		FindOneByContactTypeValue(ctx context.Context, contactType string, value string) (*UserContact, error)
		Update(ctx context.Context, data *UserContact) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserContactModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserContact struct {
		Id          int64  `db:"id"`
		UserId      int64  `db:"user_id"`
		ContactType string `db:"contact_type"`
		Value       string `db:"value"`
	}
)

func newUserContactModel(conn sqlx.SqlConn) *defaultUserContactModel {
	return &defaultUserContactModel{
		conn:  conn,
		table: "`user_contact`",
	}
}

func (m *defaultUserContactModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserContactModel) FindOne(ctx context.Context, id int64) (*UserContact, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userContactRows, m.table)
	var resp UserContact
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) FindOneByContactTypeValue(ctx context.Context, contactType string, value string) (*UserContact, error) {
	var resp UserContact
	query := fmt.Sprintf("select %s from %s where `contact_type` = ? and `value` = ? limit 1", userContactRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, contactType, value)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserContactModel) Insert(ctx context.Context, data *UserContact) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userContactRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ContactType, data.Value)
	return ret, err
}

func (m *defaultUserContactModel) Update(ctx context.Context, newData *UserContact) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userContactRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.UserId, newData.ContactType, newData.Value, newData.Id)
	return err
}

func (m *defaultUserContactModel) tableName() string {
	return m.table
}