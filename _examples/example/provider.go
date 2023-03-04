package example

import (
	"context"
	"errors"

	"github.com/getumbeluzi/providerkit"
)

func Build() *providerkit.Provider {
	exampleProvider := exampleProvider{}

	return &providerkit.Provider{
		Name:            "EXAMPLE",
		SendTransaction: providerkit.TransactionSendFunc(exampleProvider.SendTransaction),
	}
}

type exampleProvider struct{}

func (p exampleProvider) SendTransaction(ctx context.Context, provider *providerkit.Provider, transaction providerkit.Transaction) (*providerkit.Transaction, error) {
	return nil, errors.New("not implemented")
}
