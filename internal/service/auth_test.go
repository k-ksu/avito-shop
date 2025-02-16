package service

import (
	"testing"
	"time"

	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/stretchr/testify/require"
)

func TestMerch_ItemByName(t *testing.T) {
	t.Parallel()

	user := model.User{
		ID:   1,
		Name: "name",
	}

	t.Run("correct", func(t *testing.T) {
		t.Parallel()

		j := NewJWTAuth("some-key", 10*time.Minute)

		token, err := j.GenerateToken(user)
		require.NoError(t, err)

		claims, err := j.ParseToken(token)
		require.NoError(t, err)

		require.Equal(t, user.ID, claims.UserID)
		require.Equal(t, user.Name, claims.Username)
	})
}
