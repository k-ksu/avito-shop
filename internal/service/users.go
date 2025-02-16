package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/k-ksu/avito-shop/internal/errs"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UsersRepo ...
	UsersRepo interface {
		UserByName(ctx context.Context, name string) (model.User, error)
		CreateUser(ctx context.Context, name, pass string) (model.User, error)
		UsersByNames(ctx context.Context, tx wrapper.Tx, names []string) ([]model.User, error)
		LockUsers(ctx context.Context, tx wrapper.Tx, users []model.User) error
		UpdateCoins(ctx context.Context, tx wrapper.Tx, users []model.User) error
	}

	// Auther ...
	//
	//nolint:misspell
	Auther interface {
		ParseToken(tokenString string) (model.Claims, error)
		GenerateToken(user model.User) (string, error)
	}

	// Users ...
	Users struct {
		wrapper         TxWrapper
		repo            UsersRepo
		auth            Auther //nolint:misspell
		txHistorier     TxHistorier
		shopHistoryRepo ShopHistoryRepo
	}
)

// NewUsers ...
//
// nolint:misspell
func NewUsers(
	wrapper TxWrapper,
	repo UsersRepo,
	auth Auther,
	txHistorier TxHistorier,
	shopHistoryRepo ShopHistoryRepo,
) *Users {
	return &Users{
		wrapper:         wrapper,
		repo:            repo,
		auth:            auth,
		txHistorier:     txHistorier,
		shopHistoryRepo: shopHistoryRepo,
	}
}

// GetUser ...
func (u *Users) GetUser(tokenString string) (model.User, error) {
	claims, err := u.auth.ParseToken(tokenString)
	if err != nil {
		return model.User{}, fmt.Errorf("jwt.Parse: %w", err)
	}

	return model.User{
		ID:   claims.UserID,
		Name: claims.Username,
	}, nil
}

// UserInfo ...
func (u *Users) UserInfo(ctx context.Context, user model.User) (model.UserInfo, error) {
	var userInfo model.UserInfo

	err := u.wrapper.Wrap(ctx, func(tx wrapper.Tx) error {
		if err := u.repo.LockUsers(ctx, tx, []model.User{user}); err != nil {
			return fmt.Errorf("repo.LockUsers: %w", err)
		}

		users, err := u.repo.UsersByNames(ctx, tx, []string{user.Name})
		if err != nil {
			return fmt.Errorf("repo.UsersByNames: %w", err)
		}

		user = users[0]
		userInfo.Coins = user.Coins

		sent, err := u.txHistorier.GetAllFrom(ctx, tx, int(user.ID))
		if err != nil {
			return fmt.Errorf("tx.GetAllFrom: %w", err)
		}

		received, err := u.txHistorier.GetAllTo(ctx, tx, int(user.ID))
		if err != nil {
			return fmt.Errorf("tx.GetAllTo: %w", err)
		}

		userInfo.CoinHistory.Sent = sent
		userInfo.CoinHistory.Received = received

		inventory, err := u.shopHistoryRepo.GetAllByUser(ctx, tx, user.ID)
		if err != nil {
			return fmt.Errorf("shopHistoryRepo.GetAllByUser: %w", err)
		}

		userInfo.Inventory = inventory

		return nil
	})
	if err != nil {
		return model.UserInfo{}, fmt.Errorf("wrapper.Wrap: %w", err)
	}

	return userInfo, nil
}

// AuthUser ...
func (u *Users) AuthUser(ctx context.Context, name, pass string) (string, error) {
	user, err := u.repo.UserByName(ctx, name)
	if err != nil {
		if errors.Is(err, errs.ErrNoRows) {
			return u.registerUser(ctx, name, pass)
		}

		return "", fmt.Errorf("repo.UserByName: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.ObfuscatedPassword), []byte(pass)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", errs.ErrInvalidPassword
		}

		return "", fmt.Errorf("bcrypt.CompareHashAndPassword: %w: ", err)
	}

	token, err := u.auth.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("auth.GenerateToken: %w", err)
	}

	return token, nil
}

func (u *Users) registerUser(ctx context.Context, name, pass string) (string, error) {
	obfuscatedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	user, err := u.repo.CreateUser(ctx, name, string(obfuscatedPass))
	if err != nil {
		return "", fmt.Errorf("repo.CreateUser: %w", err)
	}

	token, err := u.auth.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("auth.GenerateToken: %w", err)
	}

	return token, nil
}
