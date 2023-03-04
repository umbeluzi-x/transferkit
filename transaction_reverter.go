package transferkit

import (
	"context"
)

type TransactionReverterFunc func(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error)

type TransactionReverter interface {
	RevertTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error)
}

func (t TransactionReverterFunc) RevertTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error) {
	return t(ctx, provider, transaction)
}
