package example

import (
	"context"
	"errors"

	"github.com/getumbeluzi/transferkit"
)

func Build() *transferkit.Provider {
	exampleProvider := exampleProvider{}

	return &transferkit.Provider{
		Name:            "EXAMPLE",
		SendTransaction: transferkit.TransactionSenderFunc(exampleProvider.sendTransaction),
		SupportedCurrencies: transferkit.SupportedCurrencies{
			"MZN", {}{},
		}
	}
}

type exampleProvider struct{}

func (p exampleProvider) sendTransaction(ctx context.Context, provider *transferkit.Provider, transaction transferkit.Transaction) (*transferkit.Transaction, error) {
	return nil, errors.New("not implemented")
}
