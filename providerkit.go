package providerkit

type Provider struct {
	Name            string
	SendTransaction TransactionSender
	GetTransaction  TransactionChecker
	CheckAccount    AccountChecker
	FetchAccount    AccountFetcher

	Config []byte
}
