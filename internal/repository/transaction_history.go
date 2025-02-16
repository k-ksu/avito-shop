package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
	"github.com/k-ksu/avito-shop/pkg/postgres"
)

const transactionHistoryTableName = "transaction_history"

var transactionHistoryColumns = []string{
	"from_user",
	"to_user",
	"amount",
}

// TransactionHistory ...
type TransactionHistory struct {
	*postgres.Client
}

// NewTransactionHistory ...
func NewTransactionHistory(cl *postgres.Client) *TransactionHistory {
	return &TransactionHistory{cl}
}

// AddNew ...
func (t *TransactionHistory) AddNew(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) error {
	var err error

	if tx == nil {
		tx, err = t.BeginTx(ctx, pgx.TxOptions{})
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

	qb := psql.Insert(transactionHistoryTableName).
		Columns(transactionHistoryColumns...).
		Values(transaction.FromUser, transaction.ToUser, transaction.Amount)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

// GetAllFrom ...
//
//nolint:dupl
func (t *TransactionHistory) GetAllFrom(ctx context.Context, tx wrapper.Tx, fromUser int) ([]model.SentCoins, error) {
	var err error

	if tx == nil {
		tx, err = t.BeginTx(ctx, pgx.TxOptions{})
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

	qb := psql.Select(`
		name, sum(amount) as amount
	`).From(transactionHistoryTableName).
		Join("users on to_user = users.id").
		Where(squirrel.Eq{"from_user": fromUser}).
		GroupBy("name")

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("qb.ToSql: %w", err)
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("tx.Query: %w", err)
	}

	defer rows.Close()

	var sents []model.SentCoins
	for rows.Next() {
		var sent model.SentCoins

		err = rows.Scan(&sent.ToUser, &sent.Amount)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		sents = append(sents, sent)
	}

	return sents, nil
}

// GetAllTo ...
//
//nolint:dupl
func (t *TransactionHistory) GetAllTo(ctx context.Context, tx wrapper.Tx, toUser int) ([]model.ReceivedCoins, error) {
	var err error

	if tx == nil {
		tx, err = t.BeginTx(ctx, pgx.TxOptions{})
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

	qb := psql.Select(`
		name, sum(amount) as amount
	`).From(transactionHistoryTableName).
		Join("users on from_user = users.id").
		Where(squirrel.Eq{"to_user": toUser}).
		GroupBy("name")

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("qb.ToSql: %w", err)
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("tx.Query: %w", err)
	}

	defer rows.Close()

	var sents []model.ReceivedCoins
	for rows.Next() {
		var sent model.ReceivedCoins

		err = rows.Scan(&sent.FromUser, &sent.Amount)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		sents = append(sents, sent)
	}

	return sents, nil
}
