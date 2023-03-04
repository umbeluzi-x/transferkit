package transferkit

import "context"

type TransactionCheckerFunc func(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error)

type TransactionChecker interface {
	CheckTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error)
}

func (t TransactionCheckerFunc) CheckTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error) {
	return t(ctx, provider, transaction)
}
