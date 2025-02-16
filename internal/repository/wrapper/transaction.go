package wrapper

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/k-ksu/avito-shop/pkg/postgres"
)

type (
	// Tx ...
	Tx interface {
		pgx.Tx
	}

	// Transaction ...
	Transaction struct {
		cl *postgres.Client
	}
)

// NewTransaction ...
func NewTransaction(cl *postgres.Client) *Transaction {
	return &Transaction{cl: cl}
}

// Wrap ...
func (t *Transaction) Wrap(ctx context.Context, handlers ...func(tx Tx) error) error {
	return t.cl.BeginFunc(ctx, func(tx pgx.Tx) error {
		for _, h := range handlers {
			if err := h(tx); err != nil {
				return err
			}
		}

		return nil
	})
}
