package service

import (
	"context"
	"errors"
	"testing"

	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/service/mocks"
	"github.com/stretchr/testify/require"
)

// nolint:maintidx
func TestShop_SendCoins(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name     string
		shop     *Shop
		fromUser model.User
		toUser   model.User
		amount   int32
		wantErr  bool
	}{
		{
			name: "Test 1: correct",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.AddNewMock.Expect(ctx, nil, model.Transaction{
					FromUser: 1,
					ToUser:   2,
					Amount:   100,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 900,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1100,
					},
				}).Return(nil)

				return &Shop{
					wrapper:     mockWrapper,
					txHistorier: mockTxHistorier,
					usersRepo:   mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: false,
		},
		{
			name: "Test 2: LockUsers return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(errors.New("some error"))

				return &Shop{
					wrapper:   mockWrapper,
					usersRepo: mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: true,
		},
		{
			name: "Test 3: UsersByNames return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1000,
					},
				}, errors.New("some error"))

				return &Shop{
					wrapper:   mockWrapper,
					usersRepo: mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: true,
		},
		{
			name: "Test 4: not enough users",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				return &Shop{
					wrapper:   mockWrapper,
					usersRepo: mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: true,
		},
		{
			name: "Test 5: correct with swapped users",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.AddNewMock.Expect(ctx, nil, model.Transaction{
					FromUser: 1,
					ToUser:   2,
					Amount:   100,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    2,
						Name:  "string",
						Coins: 1000,
					},
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 900,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1100,
					},
				}).Return(nil)

				return &Shop{
					wrapper:     mockWrapper,
					txHistorier: mockTxHistorier,
					usersRepo:   mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: false,
		},
		{
			name: "Test 6: not enough coins",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 50,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1000,
					},
				}, nil)

				return &Shop{
					wrapper:   mockWrapper,
					usersRepo: mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: true,
		},
		{
			name: "Test 7: UpdateCoins return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 900,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1100,
					},
				}).Return(errors.New("error"))

				return &Shop{
					wrapper:   mockWrapper,
					usersRepo: mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: true,
		},
		{
			name: "Test 8: AddNew return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.AddNewMock.Expect(ctx, nil, model.Transaction{
					FromUser: 1,
					ToUser:   2,
					Amount:   100,
				}).Return(errors.New("error"))

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
					{
						ID:   2,
						Name: "string",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
					"string",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 900,
					},
					{
						ID:    2,
						Name:  "string",
						Coins: 1100,
					},
				}).Return(nil)

				return &Shop{
					wrapper:     mockWrapper,
					txHistorier: mockTxHistorier,
					usersRepo:   mockUsersRepo,
				}
			}(),
			fromUser: model.User{
				ID:   1,
				Name: "ksu",
			},
			toUser: model.User{
				ID:   2,
				Name: "string",
			},
			amount:  100,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := tt.shop
			err := s.SendCoins(ctx, tt.fromUser, tt.toUser, tt.amount)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

// nolint:maintidx
func TestShop_BuyItem(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name    string
		shop    *Shop
		User    model.User
		item    string
		wantErr bool
	}{
		{
			name: "Test 1: correct",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 980,
					},
				}).Return(nil)

				return &Shop{
					wrapper:         mockWrapper,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: false,
		},
		{
			name: "Test 2: GetMerch return error",
			shop: func() *Shop {
				mocksMerchRepo := mocks.NewMerchRepo(t)
				mocksMerchRepo.ItemByNameMock.
					Expect(ctx, "t-shirt").
					Return(model.Merch{}, errors.New("error"))

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").
					Return(model.Merch{}, false)

				return &Shop{
					cache:     mocksMerchCacher,
					merchRepo: mocksMerchRepo,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
		{
			name: "Test 3: correct go to db",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.AddNewMock.Expect(ctx, nil, model.Transaction{
					FromUser: 1,
					ToUser:   2,
					Amount:   100,
				}).Return(nil)

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").
					Return(model.Merch{}, false)

				mocksMerchRepo := mocks.NewMerchRepo(t)
				mocksMerchRepo.ItemByNameMock.
					Expect(ctx, "t-shirt").
					Return(model.Merch{
						ID:    2,
						Name:  "t-shirt",
						Price: 20,
					}, nil)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 980,
					},
				}).Return(nil)

				return &Shop{
					wrapper:         mockWrapper,
					txHistorier:     mockTxHistorier,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
					merchRepo:       mocksMerchRepo,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: false,
		},
		{
			name: "Test 4: LockUsers return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(errors.New("error"))

				return &Shop{
					wrapper:         mockWrapper,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
		{
			name: "Test 5: UsersByNames return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{}, errors.New("error"))

				return &Shop{
					wrapper:         mockWrapper,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
		{
			name: "Test 6: UpdateCoins return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 980,
					},
				}).Return(errors.New("error"))

				return &Shop{
					wrapper:         mockWrapper,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
		{
			name: "Test 7: shopHistoryRepo return error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(errors.New("error"))

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 980,
					},
				}).Return(errors.New("error"))

				return &Shop{
					wrapper:         mockWrapper,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
		{
			name: "Test 8: not enough coins",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 10,
					},
				}, nil)

				return NewShop(mockWrapper, nil, mockUsersRepo, nil, mocksShopHistoryRepo, mocksMerchCacher)
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
		{
			name: "Test 9: AddNew returns error",
			shop: func() *Shop {
				mockWrapper := mocks.NewMockTxWrapper()

				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.GetMock.Expect("t-shirt").Return(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}, true)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(errors.New("error"))

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 1000,
					},
				}, nil)

				mockUsersRepo.UpdateCoinsMock.Expect(ctx, nil, []model.User{
					{
						ID:    1,
						Name:  "ksu",
						Coins: 980,
					},
				}).Return(nil)

				return &Shop{
					wrapper:         mockWrapper,
					usersRepo:       mockUsersRepo,
					shopHistoryRepo: mocksShopHistoryRepo,
					cache:           mocksMerchCacher,
				}
			}(),
			User: model.User{
				ID:   1,
				Name: "ksu",
			},
			item:    "t-shirt",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := tt.shop
			err := s.BuyItem(ctx, tt.User, tt.item)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestShop_WarmUpCache(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name    string
		shop    *Shop
		wantErr bool
	}{
		{
			name: "Test 1: correct",
			shop: func() *Shop {
				mocksMerchCacher := mocks.NewMerchCacher(t)
				mocksMerchCacher.AddMock.Expect(model.Merch{
					ID:    2,
					Name:  "t-shirt",
					Price: 20,
				}).Return()

				mocksMerchRepo := mocks.NewMerchRepo(t)
				mocksMerchRepo.GetAllItemsMock.Expect(ctx).Return([]model.Merch{
					{
						ID:    2,
						Name:  "t-shirt",
						Price: 20,
					},
				}, nil)

				return &Shop{
					merchRepo: mocksMerchRepo,
					cache:     mocksMerchCacher,
				}
			}(),
			wantErr: false,
		},
		{
			name: "Test 2: GetAllItems return error",
			shop: func() *Shop {
				mocksMerchRepo := mocks.NewMerchRepo(t)
				mocksMerchRepo.GetAllItemsMock.Expect(ctx).
					Return([]model.Merch{}, errors.New("error"))

				return &Shop{
					merchRepo: mocksMerchRepo,
				}
			}(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := tt.shop
			err := s.WarmUpCache(ctx)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
