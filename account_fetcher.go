package providerkit

import "context"

type AccountFetcherFunc func(ctx context.Context, provider *Provider, account Account) (*Account, error)

type AccountFetcher interface {
	FetchAccount(ctx context.Context, provider *Provider, account Account) (*Account, error)
}

func (a AccountFetcherFunc) FetchAccount(ctx context.Context, provider *Provider, account Account) (*Account, error) {
	return a(ctx, provider, account)
}
