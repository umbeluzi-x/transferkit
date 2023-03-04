package providerkit

import "context"

type TransactionCheckFunc func(ctx context.Context, provider *Provider, payment Transaction) (*Transaction, error)

type TransactionChecker interface {
	CheckTransaction(ctx context.Context, provider *Provider, payment Transaction) (*Transaction, error)
}

func TransactionGetFunc(f TransactionCheckFunc) TransactionChecker {
	return transactionChecker{
		handler: f,
	}
}

type transactionChecker struct {
	handler TransactionCheckFunc
}

func (t transactionChecker) CheckTransaction(ctx context.Context, provider *Provider, transaction Transaction) (*Transaction, error) {
	return t.handler(ctx, provider, transaction)
}
