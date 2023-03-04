package transferkit

import "context"

type BalanceFetcherFunc func(ctx context.Context, provider *Provider, account Balance) (*BalanceCheckResponse, error)

type BalanceFetcher interface {
	FetchBalance(ctx context.Context, provider *Provider, account Balance) (*BalanceCheckResponse, error)
}

func (a BalanceFetcherFunc) FetchBalance(ctx context.Context, provider *Provider, account Balance) (*BalanceCheckResponse, error) {
	return a(ctx, provider, account)
}

type BalanceCheckResponse struct{}
