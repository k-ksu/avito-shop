package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
	"github.com/k-ksu/avito-shop/pkg/postgres"
)

const shopHistoryTableName = "shop_history"

var shopHistoryColumns = []string{
	"id",
	"user_id",
	"item_id",
	"created_at",
}

// ShopHistory ...
type ShopHistory struct {
	*postgres.Client
}

// NewShopHistory ...
func NewShopHistory(cl *postgres.Client) *ShopHistory {
	return &ShopHistory{cl}
}

// AddNew ...
func (s *ShopHistory) AddNew(ctx context.Context, tx wrapper.Tx, shopHistory model.ShopHistory) error {
	var err error

	if tx == nil {
		tx, err = s.BeginTx(ctx, pgx.TxOptions{})
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

	qb := psql.Insert(shopHistoryTableName).
		Columns(shopHistoryColumns[1:3]...).
		Values(shopHistory.UserID, shopHistory.ItemID)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

// GetAllByUser ...
func (s *ShopHistory) GetAllByUser(ctx context.Context, tx wrapper.Tx, id int64) ([]model.Inventory, error) {
	var err error

	if tx == nil {
		tx, err = s.BeginTx(ctx, pgx.TxOptions{})
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

	qb := psql.Select("name", "count(*) quantity").
		From(shopHistoryTableName).
		Join("merch_items on item_id = merch_items.id").
		Where("user_id = ?", id).
		GroupBy("name")

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("qb.ToSql: %w", err)
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	defer rows.Close()

	var items []model.Inventory
	for rows.Next() {
		var item model.Inventory

		err = rows.Scan(&item.Type, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		items = append(items, item)
	}

	return items, nil
}
