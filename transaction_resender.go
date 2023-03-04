package transferkit

import (
	"context"
)

type TransactionResenderFunc func(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error)

type TransactionResender interface {
	ResendTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error)
}

func (t TransactionResenderFunc) ResendTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error) {
	return t(ctx, provider, transaction)
}
