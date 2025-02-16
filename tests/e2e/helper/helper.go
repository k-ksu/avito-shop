package helper

import (
	"context"
	"github.com/k-ksu/avito-shop/config"
	"github.com/k-ksu/avito-shop/internal/app"
	"github.com/stretchr/testify/require"
	"testing"
)

func InitUsers(t *testing.T, ctx context.Context, cfg *config.Config, usernames []string) {
	t.Helper()

	cont := app.NewContainer(ctx, cfg)

	users, err := cont.Repositories.Users.UsersByNames(ctx, nil, usernames)
	require.NoError(t, err)
	require.Len(t, users, len(usernames))

	for i := range users {
		users[i].Coins = 1000
	}

	err = cont.Repositories.Users.UpdateCoins(ctx, nil, users)
	require.NoError(t, err)
}
