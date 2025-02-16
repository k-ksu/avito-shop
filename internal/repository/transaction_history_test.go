package repository

import (
	"context"
	"testing"

	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestTransactionHistory(t *testing.T) {
	ctx := context.Background()

	t.Run("correct", func(t *testing.T) {
		users := NewUsers(testRepo)
		repo := NewTransactionHistory(testRepo)

		m1, err := users.CreateUser(ctx, "name1", "pass1")
		require.NoError(t, err)
		m2, err := users.CreateUser(ctx, "name2", "pass2")
		require.NoError(t, err)

		err = repo.AddNew(ctx, nil, model.Transaction{
			FromUser: int(m1.ID),
			ToUser:   int(m2.ID),
			Amount:   300,
		})
		require.NoError(t, err)

		sent, err := repo.GetAllFrom(ctx, nil, int(m1.ID))
		require.NoError(t, err)
		require.Len(t, sent, 1)
		require.Equal(t, int32(300), sent[0].Amount)
		require.Equal(t, m2.Name, sent[0].ToUser)

		rv, err := repo.GetAllTo(ctx, nil, int(m2.ID))
		require.NoError(t, err)
		require.Len(t, rv, 1)
		require.Equal(t, int32(300), rv[0].Amount)
		require.Equal(t, m1.Name, rv[0].FromUser)
	})
}
