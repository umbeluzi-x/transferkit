package transferkit

import "context"

type TransactionCheckerFunc func(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionCheckResponse, error)

type TransactionChecker interface {
	CheckTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionCheckResponse, error)
}

func (t TransactionCheckerFunc) CheckTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionCheckResponse, error) {
	return t(ctx, provider, transaction)
}

type TransactionCheckResponse struct{}
