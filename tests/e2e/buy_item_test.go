package e2e

import (
	"context"
	"github.com/k-ksu/avito-shop/config"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/tests/e2e/helper"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	testUserName1 = "test1"
	testUserPass1 = "pass1"

	testUserName2 = "test2"
	testUserPass2 = "pass2"
)

func TestBuyItem(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cfg, err := config.New()
		require.NoError(t, err)

		cfg.URL = cfg.URLLocal

		addr := "http://" + "127.0.0.1:" + cfg.Port
		token1, err := helper.AuthUser(addr, testUserName1, testUserPass1)
		require.NoError(t, err)

		helper.InitUsers(t, ctx, cfg, []string{testUserName1})

		userInfoPrev := helper.UserInfo(t, addr, token1)

		helper.BuyItem(t, "cup", addr, token1)

		userInfoCur := helper.UserInfo(t, addr, token1)

		// смотрим что списались монетки
		require.Equal(t, userInfoPrev.Coins-20, userInfoCur.Coins)
		// смотрим что кружка добавилась
		require.Equal(t, getItemCount(userInfoPrev, "cup")+1, getItemCount(userInfoCur, "cup"))
	})
}

func getItemCount(userInfo model.UserInfo, item string) int {
	for _, inventory := range userInfo.Inventory {
		if inventory.Type == item {
			return int(inventory.Quantity)
		}
	}

	return 0
}
