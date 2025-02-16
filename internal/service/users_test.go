package service

import (
	"context"
	"errors"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/service/mocks"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestUsers_GetUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name    string
		users   *Users
		token   string
		wantErr bool
	}{
		{
			name: "Test 1: correct",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.AddNewMock.Expect(ctx, nil, model.Transaction{
					FromUser: 1,
					ToUser:   2,
					Amount:   100,
				}).Return(nil)

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

				mockAuther := mocks.NewAuther(t)
				mockAuther.ParseTokenMock.Expect("someToken").
					Return(model.Claims{
						UserID:         1,
						Username:       "ksu",
						StandardClaims: jwt.StandardClaims{},
					}, nil)

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

				return &Users{
					wrapper:         mockWrapper,
					shopHistoryRepo: mocksShopHistoryRepo,
					auth:            mockAuther, //nolint:misspell
					txHistorier:     mockTxHistorier,
					repo:            mockUsersRepo,
				}
			}(),
			token:   "someToken",
			wantErr: false,
		},
		{
			name: "Test 2: ParseToken return error",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.AddNewMock.Expect(ctx, nil, model.Transaction{
					FromUser: 1,
					ToUser:   2,
					Amount:   100,
				}).Return(nil)

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

				mockAuther := mocks.NewAuther(t)
				mockAuther.ParseTokenMock.Expect("someToken").
					Return(model.Claims{}, errors.New("some error"))

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

				return &Users{
					wrapper:         mockWrapper,
					shopHistoryRepo: mocksShopHistoryRepo,
					auth:            mockAuther, //nolint:misspell
					txHistorier:     mockTxHistorier,
					repo:            mockUsersRepo,
				}
			}(),
			token:   "someToken",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := tt.users
			_, err := s.GetUser(tt.token)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

// nolint:maintidx
func TestUsers_UserInfo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name    string
		users   *Users
		user    model.User
		wantErr bool
	}{
		{
			name: "Test 1: correct",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.GetAllFromMock.Expect(ctx, nil, 1).
					Return([]model.SentCoins{
						{
							Amount: 400,
							ToUser: "user1",
						},
					}, nil)

				mockTxHistorier.GetAllToMock.Expect(ctx, nil, 1).
					Return([]model.ReceivedCoins{
						{
							Amount:   50,
							FromUser: "user1",
						},
					}, nil)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mocksShopHistoryRepo.GetAllByUserMock.Expect(ctx, nil, 1).
					Return([]model.Inventory{
						{
							Quantity: 2,
							Type:     "cup",
						},
					}, nil)

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

				mockAuther := mocks.NewAuther(t)
				mockAuther.ParseTokenMock.Expect("someToken").
					Return(model.Claims{
						UserID:         1,
						Username:       "ksu",
						StandardClaims: jwt.StandardClaims{},
					}, nil)

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

				return &Users{
					wrapper:         mockWrapper,
					shopHistoryRepo: mocksShopHistoryRepo,
					auth:            mockAuther, //nolint:misspell
					txHistorier:     mockTxHistorier,
					repo:            mockUsersRepo,
				}
			}(),
			user: model.User{
				ID:   1,
				Name: "ksu",
			},
			wantErr: false,
		},
		{
			name: "Test 2: LockUsers return error",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).
					Return(errors.New("some error"))

				return &Users{
					wrapper: mockWrapper,
					repo:    mockUsersRepo,
				}
			}(),
			user: model.User{
				ID:   1,
				Name: "ksu",
			},
			wantErr: true,
		},
		{
			name: "Test 3: UsersByNames return error",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.LockUsersMock.Expect(ctx, nil, []model.User{
					{
						ID:   1,
						Name: "ksu",
					},
				}).Return(nil)

				mockUsersRepo.UsersByNamesMock.Expect(ctx, nil, []string{
					"ksu",
				}).Return([]model.User{}, errors.New("some error"))

				return &Users{
					wrapper: mockWrapper,
					repo:    mockUsersRepo,
				}
			}(),
			user: model.User{
				ID:   1,
				Name: "ksu",
			},
			wantErr: true,
		},
		{
			name: "Test 4: GetAllFrom return error",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.GetAllFromMock.Expect(ctx, nil, 1).
					Return([]model.SentCoins{}, errors.New("some error"))

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

				mockAuther := mocks.NewAuther(t)
				mockAuther.ParseTokenMock.Expect("someToken").
					Return(model.Claims{
						UserID:         1,
						Username:       "ksu",
						StandardClaims: jwt.StandardClaims{},
					}, nil)

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

				return &Users{
					wrapper:         mockWrapper,
					shopHistoryRepo: mocksShopHistoryRepo,
					auth:            mockAuther, //nolint:misspell
					txHistorier:     mockTxHistorier,
					repo:            mockUsersRepo,
				}
			}(),
			user: model.User{
				ID:   1,
				Name: "ksu",
			},
			wantErr: true,
		},
		{
			name: "Test 5: GetAllTo return error",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.GetAllFromMock.Expect(ctx, nil, 1).
					Return([]model.SentCoins{
						{
							Amount: 400,
							ToUser: "user1",
						},
					}, nil)

				mockTxHistorier.GetAllToMock.Expect(ctx, nil, 1).
					Return([]model.ReceivedCoins{}, errors.New("some error"))

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

				mockAuther := mocks.NewAuther(t)
				mockAuther.ParseTokenMock.Expect("someToken").
					Return(model.Claims{
						UserID:         1,
						Username:       "ksu",
						StandardClaims: jwt.StandardClaims{},
					}, nil)

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

				return &Users{
					wrapper:         mockWrapper,
					shopHistoryRepo: mocksShopHistoryRepo,
					auth:            mockAuther, //nolint:misspell
					txHistorier:     mockTxHistorier,
					repo:            mockUsersRepo,
				}
			}(),
			user: model.User{
				ID:   1,
				Name: "ksu",
			},
			wantErr: true,
		},
		{
			name: "Test 6: GetAllByUser return error",
			users: func() *Users {
				mockWrapper := mocks.NewMockTxWrapper()

				mockTxHistorier := mocks.NewTxHistorier(t)
				mockTxHistorier.GetAllFromMock.Expect(ctx, nil, 1).
					Return([]model.SentCoins{
						{
							Amount: 400,
							ToUser: "user1",
						},
					}, nil)

				mockTxHistorier.GetAllToMock.Expect(ctx, nil, 1).
					Return([]model.ReceivedCoins{
						{
							Amount:   50,
							FromUser: "user1",
						},
					}, nil)

				mocksShopHistoryRepo := mocks.NewShopHistoryRepo(t)
				mocksShopHistoryRepo.AddNewMock.Expect(ctx, nil, model.ShopHistory{
					UserID: 1,
					ItemID: 2,
				}).Return(nil)

				mocksShopHistoryRepo.GetAllByUserMock.Expect(ctx, nil, 1).
					Return([]model.Inventory{}, errors.New("some error"))

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

				mockAuther := mocks.NewAuther(t)
				mockAuther.ParseTokenMock.Expect("someToken").
					Return(model.Claims{
						UserID:         1,
						Username:       "ksu",
						StandardClaims: jwt.StandardClaims{},
					}, nil)

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

				return NewUsers(mockWrapper, mockUsersRepo, mockAuther, mockTxHistorier, mocksShopHistoryRepo)
			}(),
			user: model.User{
				ID:   1,
				Name: "ksu",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := tt.users
			_, err := s.UserInfo(ctx, tt.user)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUsers_AuthUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	obfuscatedPass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	notCorrectPass, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)

	tests := []struct {
		name     string
		users    *Users
		userName string
		password string
		wantErr  bool
	}{
		{
			name: "Test 1: correct",
			users: func() *Users {
				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.UserByNameMock.Expect(ctx, "ksu").
					Return(model.User{
						ID:                 1,
						Name:               "ksu",
						ObfuscatedPassword: string(obfuscatedPass),
						Coins:              1000,
					}, nil)

				mockAuther := mocks.NewAuther(t)
				mockAuther.GenerateTokenMock.Expect(model.User{
					ID:                 1,
					Name:               "ksu",
					ObfuscatedPassword: string(obfuscatedPass),
					Coins:              1000,
				}).
					Return("correctToken", nil)

				return &Users{
					auth: mockAuther, //nolint:misspell
					repo: mockUsersRepo,
				}
			}(),
			userName: "ksu",
			password: "password",
			wantErr:  false,
		},
		{
			name: "Test 2: incorrect password",
			users: func() *Users {
				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.UserByNameMock.Expect(ctx, "ksu").
					Return(model.User{
						ID:                 1,
						Name:               "ksu",
						ObfuscatedPassword: string(notCorrectPass),
						Coins:              1000,
					}, nil)

				return &Users{
					repo: mockUsersRepo,
				}
			}(),
			userName: "ksu",
			password: "password",
			wantErr:  true,
		},
		{
			name: "Test 3: UserByName return error",
			users: func() *Users {
				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.UserByNameMock.Expect(ctx, "ksu").
					Return(model.User{}, errors.New("some error"))

				return &Users{
					repo: mockUsersRepo,
				}
			}(),
			userName: "ksu",
			password: "password",
			wantErr:  true,
		},
		{
			name: "Test 4: correct",
			users: func() *Users {
				mockUsersRepo := mocks.NewUsersRepo(t)
				mockUsersRepo.UserByNameMock.Expect(ctx, "ksu").
					Return(model.User{
						ID:                 1,
						Name:               "ksu",
						ObfuscatedPassword: string(obfuscatedPass),
						Coins:              1000,
					}, nil)

				mockAuther := mocks.NewAuther(t)
				mockAuther.GenerateTokenMock.Expect(model.User{
					ID:                 1,
					Name:               "ksu",
					ObfuscatedPassword: string(obfuscatedPass),
					Coins:              1000,
				}).
					Return("", errors.New("some error"))

				return &Users{
					auth: mockAuther, //nolint:misspell
					repo: mockUsersRepo,
				}
			}(),
			userName: "ksu",
			password: "password",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := tt.users
			_, err := s.AuthUser(ctx, tt.userName, tt.password)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
