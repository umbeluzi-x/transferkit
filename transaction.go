package transferkit

type TransactionStatus string

const (
	TransactionStatusSuccess    = TransactionStatus("SUCCESS")
	TransactionStatusUnknown    = TransactionStatus("UNKNOWN")
	TransactionStatusProcessing = TransactionStatus("PROCESSING")
	TransactionStatusFailure    = TransactionStatus("FAILURE")
)

type TransactionFailureReason string

const (
	TransactionFailureNotEnoughBalance = TransactionFailureReason("NOT_ENOUGH_BALANCE")
	TransactionFailureInvalidAccount   = TransactionFailureReason("INVALID_ACCOUNT")
)

type Transaction struct {
	ID             int64             `json:"id"`
	Status         TransactionStatus `json:"status"`
	DeliveryMethod DeliveryMethod    `json:"delivery_method"`
	Amount         Money             `json:"amount"`
}

type Money struct {
	Value    float64  `json:"value"`
	Currency Currency `json:"currency"`
}
