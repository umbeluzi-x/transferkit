package providerkit

import "context"

type BalanceFetcherFunc func(ctx context.Context, provider *Provider, account Balance) (*Balance, error)

type BalanceFetcher interface {
	FetchBalance(ctx context.Context, provider *Provider, account Balance) (*Balance, error)
}

func (a BalanceFetcherFunc) FetchBalance(ctx context.Context, provider *Provider, account Balance) (*Balance, error) {
	return a(ctx, provider, account)
}
