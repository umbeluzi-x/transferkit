package example

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/getumbeluzi/transferkit"
)

func Build() *transferkit.Provider {
	exampleProvider := exampleProvider{}

	return &transferkit.Provider{
		Name:    "EXAMPLE",
		Version: "0.1.0",
		Config:  []byte(`{
			"type": "object",
			"properties": {
				"base_url": {
					"type": "string"
				}
			}
		}`)
		OnInit: transferkit.InitterFunc(exampleProvider.init),
		OnTransactionSend: transferkit.TransactionSenderFunc(exampleProvider.sendTransaction),
	}
}

type exampleProvider struct{
	config *struct{}
}

func (p *exampleProvider) init(ctx context.Context, *opts transferkit.ProviderOptions) error {
	if err := json.Unmarshal(opts, &p.config); err != nil {
		return err
	}

	return nil
}

func (p exampleProvider) sendTransaction(ctx context.Context, transaction transferkit.Transaction) (*transferkit.Transaction, error) {
	return nil, errors.New("not implemented")
}
