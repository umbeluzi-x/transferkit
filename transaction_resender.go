package transferkit

import (
	"context"
)

type TransactionResenderFunc func(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionResendResponse, error)

type TransactionResender interface {
	ResendTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionResendResponse, error)
}

func (t TransactionResenderFunc) ResendTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionResendResponse, error) {
	return t(ctx, provider, transaction)
}

type TransactionResendResponse struct{}
