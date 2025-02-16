package mocks

import (
	"context"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
)

type MockTxWrapper struct{}

func NewMockTxWrapper() *MockTxWrapper {
	return &MockTxWrapper{}
}

func (t *MockTxWrapper) Wrap(_ context.Context, handlers ...func(tx wrapper.Tx) error) error {
	for _, handler := range handlers {
		if err := handler(nil); err != nil {
			return err
		}
	}

	return nil
}
