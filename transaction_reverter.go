package transferkit

import (
	"context"
)

type TransactionReverterFunc func(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionRevertResponse, error)

type TransactionReverter interface {
	RevertTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionRevertResponse, error)
}

func (t TransactionReverterFunc) RevertTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*TransactionRevertResponse, error) {
	return t(ctx, provider, transaction)
}

type TransactionRevertResponse struct {
	Status TransactionStatus `json:"status"`
}
