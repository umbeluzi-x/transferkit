package providerkit

import "context"

type TransactionSenderFunc func(ctx context.Context, provider *Provider, payment Transaction) (*Transaction, error)

type TransactionSender interface {
	SendTransaction(ctx context.Context, provider *Provider, payment Transaction) (*Transaction, error)
}

func (t TransactionSenderFunc) SendTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error) {
	return t(ctx, provider, transaction)
}
