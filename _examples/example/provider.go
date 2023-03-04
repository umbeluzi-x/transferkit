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
		SendTransaction: providerkit.TransactionSenderFunc(exampleProvider.sendTransaction),
		SupportedCurrencies: providerkit.SupportedCurrencies{
			"MZN", {}{},
		}
	}
}

type exampleProvider struct{}

func (p exampleProvider) sendTransaction(ctx context.Context, provider *providerkit.Provider, transaction providerkit.Transaction) (*providerkit.Transaction, error) {
	return nil, errors.New("not implemented")
}
