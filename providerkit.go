package providerkit

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

type PaymentMethod string

const (
	PaymentMethodMobileMoney = PaymentMethod("MOBILE_MONEY")
)
