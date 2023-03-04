package providerkit

import "context"

type TransactionSenderFunc func(ctx context.Context, provider *Provider, payment Transaction) (*Transaction, error)

type TransactionSender interface {
	SendTransaction(ctx context.Context, provider *Provider, payment Transaction) (*Transaction, error)
}

func TransactionSendFunc(f TransactionSenderFunc) TransactionSender {
	return transactionSender{
		handler: f,
	}
}

type transactionSender struct {
	handler TransactionSenderFunc
}

func (t transactionSender) SendTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error) {
	return t.handler(ctx, provider, transaction)
}
