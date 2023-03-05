package transferkit

type DeliveryMethodType string

const (
	DeliveryMethodTypeMobileTransfer = DeliveryMethodType("MOBILE_TRANSFER")
	DeliveryMethodTypeBankTransfer   = DeliveryMethodType("BANK_TRANSFER")
)

type DeliveryMethod struct {
	Type          DeliveryMethodType `json:"type"`
	MSISDN        *string            `json:"msisdn"`
	AccountNumber *string            `json:"account_number"`
}
