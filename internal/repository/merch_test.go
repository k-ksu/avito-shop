package repository

import (
	"context"
	"testing"

	"github.com/k-ksu/avito-shop/internal/errs"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest,tparallel
func TestMerch_ItemByName(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("correct", func(t *testing.T) {
		repo := NewMerch(testRepo)

		m, err := repo.ItemByName(ctx, "cup")
		require.NoError(t, err)
		require.Equal(t, int32(20), m.Price)
	})

	t.Run("unknown", func(t *testing.T) {
		repo := NewMerch(testRepo)

		_, err := repo.ItemByName(ctx, "cup2")
		require.Equal(t, err, errs.ErrNoRows)
	})
}

// nolint:tparallel
func TestMerch_GetAllItems(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("correct", func(t *testing.T) {
		repo := NewMerch(testRepo)

		m, err := repo.GetAllItems(ctx)
		require.NoError(t, err)
		require.Len(t, m, 10)
	})
}
