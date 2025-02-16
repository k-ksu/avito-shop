package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/k-ksu/avito-shop/internal/errs"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/pkg/postgres"
)

const merchItemsTableName = "merch_items"

var merchItemsColumns = []string{
	"id",
	"name",
	"price",
}

type (
	// Merch ...
	Merch struct {
		*postgres.Client
	}
)

// NewMerch ...
func NewMerch(cl *postgres.Client) *Merch {
	return &Merch{cl}
}

// ItemByName ...
func (m *Merch) ItemByName(ctx context.Context, name string) (model.Merch, error) {
	qb := psql.Select(merchItemsColumns...).
		From(merchItemsTableName).
		Where(squirrel.Eq{"name": name})

	sql, args, err := qb.ToSql()
	if err != nil {
		return model.Merch{}, fmt.Errorf("qb.ToSql: %w", err)
	}

	var merch model.Merch
	if err = m.QueryRow(ctx, sql, args...).Scan(&merch.ID, &merch.Name, &merch.Price); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Merch{}, errs.ErrNoRows
		}

		return model.Merch{}, fmt.Errorf("QueryRow: %w", err)
	}

	return merch, nil
}

// GetAllItems ...
func (m *Merch) GetAllItems(ctx context.Context) ([]model.Merch, error) {
	qb := psql.Select(merchItemsColumns...).From(merchItemsTableName)
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("psql.Select: %w", err)
	}

	rows, err := m.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("m.Query: %w", err)
	}

	defer rows.Close()

	var merchs []model.Merch
	for rows.Next() {
		var merch model.Merch
		if err = rows.Scan(&merch.ID, &merch.Name, &merch.Price); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		merchs = append(merchs, merch)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return merchs, nil
}
