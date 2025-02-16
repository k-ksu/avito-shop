package e2e

import (
	"context"
	"github.com/k-ksu/avito-shop/config"
	"github.com/k-ksu/avito-shop/tests/e2e/helper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendCoin(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cfg, err := config.New()
		require.NoError(t, err)

		cfg.URL = cfg.URLLocal

		addr := "http://" + cfg.Host + ":" + cfg.Port
		_, err = helper.AuthUser(addr, testUserName1, testUserPass1)
		require.NoError(t, err)
		token2, err := helper.AuthUser(addr, testUserName2, testUserPass2)
		require.NoError(t, err)

		helper.InitUsers(t, ctx, cfg, []string{testUserName1, testUserName2})

		userInfoPrev2 := helper.UserInfo(t, addr, token2)

		helper.SendCoin(t, addr, testUserName1, token2)

		userInfoCur2 := helper.UserInfo(t, addr, token2)

		// смотрим что списались монетки
		require.Equal(t, userInfoPrev2.Coins-100, userInfoCur2.Coins)
	})
}
