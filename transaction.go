package providerkit

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
	ID     int64             `json:"id"`
	Status TransactionStatus `json:"status"`
}
