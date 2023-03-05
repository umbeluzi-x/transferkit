package transferkit

import (
	"context"
	"errors"
)

type TransactionSenderMiddleware func(TransactionSender) TransactionSender

func RestrictCurrency(currency Currency, handler TransactionSender) TransactionSender {
	return TransactionSenderFunc(func(ctx context.Context, transaction *Transaction) (*TransactionSenderResponse, error) {
		if transaction.DeliveryAmount.Currency != currency {
			return nil, errors.New("invalid currency")
		}

		return handler.SendTransaction(ctx, transaction)
	})
}

func RestrictAmount(min, max float64) TransactionSenderMiddleware {
	return func(ts TransactionSender) TransactionSender {
		return TransactionSenderFunc(func(ctx context.Context, transaction *Transaction) (*TransactionSenderResponse, error) {
			if transaction.DeliveryAmount.Value < min || transaction.DeliveryAmount.Value > max {
				return nil, errors.New("amount out of bounds")
			}

			return ts.SendTransaction(ctx, transaction)
		})
	}
}
