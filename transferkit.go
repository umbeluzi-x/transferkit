package transferkit

import (
	"errors"
)

var (
	ErrInvalidDeliveryMethodType = errors.New("invalid payment method")
	ErrInvalidCurrency           = errors.New("invalid currency")
)

type Provider struct {
	Name   string
	Config []byte
	Operations
	Requirements
}

type Requirements struct {
	SupportedCurrencies     []Currency
	SupportedPaymentMethods []DeliveryMethodType
}

type Operations struct {
	SendTransaction TransactionSender
	GetTransaction  TransactionChecker
	CheckAccount    AccountChecker
	FetchAccount    AccountFetcher
	FetchBalance    BalanceFetcher
}

type Currency string

const (
	CurrencyMZN = "MZN"
)

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
