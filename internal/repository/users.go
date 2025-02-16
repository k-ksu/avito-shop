package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/k-ksu/avito-shop/internal/errs"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
	"github.com/k-ksu/avito-shop/pkg/postgres"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var (
	usersTableName = "users"

	usersTableColumns = []string{
		"id",
		"name",
		"obfuscated_password",
		"coins",
	}
)

// Users ...
type Users struct {
	*postgres.Client
}

// NewUsers ...
func NewUsers(cl *postgres.Client) *Users {
	return &Users{cl}
}

// UserByName ...
func (u *Users) UserByName(ctx context.Context, name string) (model.User, error) {
	qb := psql.Select(usersTableColumns...).
		From(usersTableName).
		Where(squirrel.Eq{"name": name})

	sql, args, err := qb.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("qb.ToSql: %w", err)
	}

	var user model.User
	if err = u.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Name, &user.ObfuscatedPassword, &user.Coins); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, errs.ErrNoRows
		}

		return model.User{}, fmt.Errorf("QueryRow: %w", err)
	}

	return user, nil
}

// UsersByNames ...
func (u *Users) UsersByNames(ctx context.Context, tx wrapper.Tx, names []string) ([]model.User, error) {
	var err error

	if tx == nil {
		tx, err = u.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return nil, fmt.Errorf("BeginTx: %w", err)
		}

		defer func() {
			if err != nil {
				tx.Rollback(ctx) //nolint:errcheck

				return
			}

			tx.Commit(ctx) //nolint:errcheck
		}()
	}

	qb := psql.Select(usersTableColumns...).
		From(usersTableName).
		Where(squirrel.Eq{"name": names})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("qb.ToSql: %w", err)
	}

	var users []model.User

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNoRows
		}

		return nil, fmt.Errorf("QueryRow: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Name, &user.ObfuscatedPassword, &user.Coins); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

// CreateUser ...
func (u *Users) CreateUser(ctx context.Context, name, pass string) (model.User, error) {
	qb := psql.Insert(usersTableName).
		Columns(usersTableColumns[1:3]...).
		Values(name, pass).
		Suffix("returning id")

	sql, args, err := qb.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("qb.ToSql: %w", err)
	}

	var id int64
	if err = u.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return model.User{}, fmt.Errorf("QueryRow: %w", err)
	}

	return model.User{
		ID:                 id,
		Name:               name,
		ObfuscatedPassword: pass,
	}, nil
}

// LockUsers ...
func (u *Users) LockUsers(ctx context.Context, tx wrapper.Tx, users []model.User) error {
	if len(users) == 0 {
		return nil
	}

	var err error

	if tx == nil {
		tx, err = u.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return fmt.Errorf("BeginTx: %w", err)
		}

		defer func() {
			if err != nil {
				tx.Rollback(ctx) //nolint:errcheck

				return
			}

			tx.Commit(ctx) //nolint:errcheck
		}()
	}

	locks := make([]string, 0, len(users))
	for _, user := range users {
		locks = append(locks, fmt.Sprintf("pg_advisory_xact_lock(%d)", user.ID))
	}

	qb := psql.Select(locks...).From(usersTableName)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

// UpdateCoins ...
func (u *Users) UpdateCoins(ctx context.Context, tx wrapper.Tx, users []model.User) error {
	if len(users) == 0 {
		return nil
	}

	var err error

	if tx == nil {
		tx, err = u.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return fmt.Errorf("BeginTx: %w", err)
		}

		defer func() {
			if err != nil {
				tx.Rollback(ctx) //nolint:errcheck

				return
			}

			tx.Commit(ctx) //nolint:errcheck
		}()
	}

	qb := psql.Insert(usersTableName).Columns("id", "coins")
	for _, user := range users {
		qb = qb.Values(user.ID, user.Coins)
	}

	qb = qb.Suffix("on conflict(id) do update set coins = EXCLUDED.coins")
	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
