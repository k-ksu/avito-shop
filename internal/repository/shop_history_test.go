package repository

import (
	"context"
	"testing"

	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestShopHistory(t *testing.T) {
	ctx := context.Background()

	t.Run("correct", func(t *testing.T) {
		users := NewUsers(testRepo)
		repo := NewShopHistory(testRepo)

		m12, err := users.CreateUser(ctx, "name12", "pass12")
		require.NoError(t, err)

		err = repo.AddNew(ctx, nil, model.ShopHistory{
			UserID: m12.ID,
			ItemID: 1,
		})
		require.NoError(t, err)

		inventory, err := repo.GetAllByUser(ctx, nil, m12.ID)
		require.NoError(t, err)
		require.Len(t, inventory, 1)
		require.Equal(t, int32(1), inventory[0].Quantity)
		require.Equal(t, "t-shirt", inventory[0].Type)
	})
}
