package service

import (
	"context"
	"fmt"

	"github.com/k-ksu/avito-shop/internal/errs"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
)

type (
	// MerchCacher ...
	MerchCacher interface {
		Add(item model.Merch)
		Get(name string) (model.Merch, bool)
	}

	// TxWrapper ...
	TxWrapper interface {
		Wrap(ctx context.Context, handlers ...func(tx wrapper.Tx) error) error
	}

	// TxHistorier ...
	TxHistorier interface {
		AddNew(ctx context.Context, tx wrapper.Tx, transaction model.Transaction) error
		GetAllFrom(ctx context.Context, tx wrapper.Tx, fromUser int) ([]model.SentCoins, error)
		GetAllTo(ctx context.Context, tx wrapper.Tx, toUser int) ([]model.ReceivedCoins, error)
	}

	// MerchRepo ...
	MerchRepo interface {
		ItemByName(ctx context.Context, name string) (model.Merch, error)
		GetAllItems(ctx context.Context) ([]model.Merch, error)
	}

	// ShopHistoryRepo ...
	ShopHistoryRepo interface {
		AddNew(ctx context.Context, tx wrapper.Tx, shopHistory model.ShopHistory) error
		GetAllByUser(ctx context.Context, tx wrapper.Tx, id int64) ([]model.Inventory, error)
	}

	// Shop ...
	Shop struct {
		wrapper         TxWrapper
		txHistorier     TxHistorier
		usersRepo       UsersRepo
		merchRepo       MerchRepo
		shopHistoryRepo ShopHistoryRepo
		cache           MerchCacher
	}
)

// NewShop ...
func NewShop(
	wrapper TxWrapper,
	txHistorier TxHistorier,
	usersRepo UsersRepo,
	merchRepo MerchRepo,
	shopHistoryRepo ShopHistoryRepo,
	cache MerchCacher,
) *Shop {
	return &Shop{
		wrapper:         wrapper,
		usersRepo:       usersRepo,
		txHistorier:     txHistorier,
		merchRepo:       merchRepo,
		shopHistoryRepo: shopHistoryRepo,
		cache:           cache,
	}
}

// SendCoins ...
func (s *Shop) SendCoins(ctx context.Context, fromUser, toUser model.User, amount int32) error {
	return s.wrapper.Wrap(ctx, func(tx wrapper.Tx) error {
		if err := s.usersRepo.LockUsers(ctx, tx, []model.User{fromUser, toUser}); err != nil {
			return fmt.Errorf("usersRepo.LockUsers: %w", err)
		}

		users, err := s.usersRepo.UsersByNames(ctx, tx, []string{fromUser.Name, toUser.Name})
		if err != nil {
			return fmt.Errorf("usersRepo.UsersByNames: %w", err)
		}

		//nolint:mnd
		if len(users) != 2 { // предполагаем что fromUser был проверен на этапе раскодировки токена
			return errs.ErrUserNotExists
		}

		if users[0].Name != fromUser.Name { // fromUser первый
			users[0], users[1] = users[1], users[0]
		}

		fromUser = users[0]
		toUser = users[1]

		if amount > fromUser.Coins {
			return errs.ErrNotEnoughMoney
		}

		fromUser.Coins -= amount
		toUser.Coins += amount
		if err = s.usersRepo.UpdateCoins(ctx, tx, []model.User{fromUser, toUser}); err != nil {
			return fmt.Errorf("usersRepo.UpdateCoins: %w", err)
		}

		if err = s.txHistorier.AddNew(ctx, tx, model.Transaction{
			FromUser: int(fromUser.ID),
			ToUser:   int(toUser.ID),
			Amount:   amount,
		}); err != nil {
			return fmt.Errorf("txHistorier.AddNew: %w", err)
		}

		return nil
	})
}

// BuyItem ...
func (s *Shop) BuyItem(ctx context.Context, user model.User, item string) error {
	merch, err := s.GetMerch(ctx, item)
	if err != nil {
		return fmt.Errorf("s.GetMerch: %w", err)
	}

	return s.wrapper.Wrap(ctx, func(tx wrapper.Tx) error {
		if err = s.usersRepo.LockUsers(ctx, tx, []model.User{user}); err != nil {
			return fmt.Errorf("usersRepo.LockUsers: %w", err)
		}

		users, err := s.usersRepo.UsersByNames(ctx, tx, []string{user.Name})
		if err != nil {
			return fmt.Errorf("usersRepo.UsersByNames: %w", err)
		}

		// на данном этапе как минимум один пользователь есть так как
		// в случае с нулем пользователем выйдем от ошибки выше
		user = users[0]
		if merch.Price > user.Coins {
			return errs.ErrNotEnoughMoney
		}

		user.Coins -= merch.Price
		if err = s.usersRepo.UpdateCoins(ctx, tx, []model.User{user}); err != nil {
			return fmt.Errorf("usersRepo.UpdateCoins: %w", err)
		}

		if err = s.shopHistoryRepo.AddNew(ctx, tx, model.ShopHistory{
			UserID: user.ID,
			ItemID: merch.ID,
		}); err != nil {
			return fmt.Errorf("shopHistoryRepo.AddNew: %w", err)
		}

		return nil
	})
}

// GetMerch ...
func (s *Shop) GetMerch(ctx context.Context, item string) (model.Merch, error) {
	merch, ok := s.cache.Get(item)
	if ok {
		return merch, nil
	}

	merch, err := s.merchRepo.ItemByName(ctx, item)
	if err != nil {
		return model.Merch{}, fmt.Errorf("merchRepo.ItemByName: %w", err)
	}

	return merch, nil
}

// WarmUpCache ...
func (s *Shop) WarmUpCache(ctx context.Context) error {
	merchs, err := s.merchRepo.GetAllItems(ctx)
	if err != nil {
		return fmt.Errorf("m.GetAllItems: %w", err)
	}

	for _, merch := range merchs {
		s.cache.Add(merch)
	}

	return nil
}
