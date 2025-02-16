package repository

import (
	"context"
	"testing"

	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/stretchr/testify/require"
)

// nolint:paralleltest
func TestUsers(t *testing.T) {
	ctx := context.Background()

	t.Run("correct", func(t *testing.T) {
		repo := NewUsers(testRepo)

		m, err := repo.CreateUser(ctx, "name", "pass")
		require.NoError(t, err)

		us, err := repo.UserByName(ctx, "name")
		require.NoError(t, err)
		require.Equal(t, int32(1000), us.Coins)

		uss, err := repo.UsersByNames(ctx, nil, []string{"name"})
		require.NoError(t, err)
		require.Len(t, uss, 1)

		require.Equal(t, us, uss[0])

		us.Coins -= 100
		err = repo.UpdateCoins(ctx, nil, []model.User{us})
		require.NoError(t, err)

		us, err = repo.UserByName(ctx, "name")
		require.NoError(t, err)
		require.Equal(t, int32(900), us.Coins)

		_, err = testRepo.Exec(ctx, `
			delete from users where id = $1
		`, m.ID)
		require.NoError(t, err)
	})
}
