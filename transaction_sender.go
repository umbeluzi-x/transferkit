package transferkit

import (
	"context"
)

type TransactionSender interface {
	SendTransaction(ctx context.Context, transaction *Transaction) (*TransactionSenderResponse, error)
}

type TransactionSenderFunc func(ctx context.Context, transaction *Transaction) (*TransactionSenderResponse, error)

func (t TransactionSenderFunc) SendTransaction(ctx context.Context, transaction *Transaction) (*TransactionSenderResponse, error) {
	return t(ctx, transaction)
}

type TransactionSenderResponse struct {
	Status TransactionStatus        `json:"status"`
	Reason TransactionFailureReason `json:"reason"`
}
