package transferkit

import "context"

type AccountFetcherFunc func(ctx context.Context, provider *Provider, account AccountFetch) (*AccountFetchResponse, error)

type AccountFetcher interface {
	FetchAccount(ctx context.Context, provider *Provider, account AccountFetch) (*AccountFetchResponse, error)
}

func (a AccountFetcherFunc) FetchAccount(ctx context.Context, provider *Provider, account AccountFetch) (*AccountFetchResponse, error) {
	return a(ctx, provider, account)
}

type AccountFetchResponse struct{}
