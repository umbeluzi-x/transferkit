package transferkit

import (
	"errors"
)

var (
	ErrInvalidPaymentMethod = errors.New("invalid payment method")
)

type Provider struct {
	Name   string
	Config []byte
	Operations
	Requirements
}

type Requirements struct {
	SupportedCurrencies     []Currency
	SupportedPaymentMethods []PaymentMethod
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

type DeliveryMethod string

const (
	DeliveryMethodMobileTransfer = DeliveryMethod("MOBILE_TRANSFER")
	DeliveryMethodBankTransfer   = DeliveryMethod("BANK_TRANSFER")
)
