package providerkit

import "context"

type AccountFetcherFunc func(ctx context.Context, provider *Provider, account Account) (*Account, error)

type AccountFetcher interface {
	FetchAccount(ctx context.Context, provider *Provider, account Account) (*Account, error)
}

func AccountFetchFunc(f AccountFetcherFunc) AccountFetcher {
	return accountFetcher{
		handler: f,
	}
}

type accountFetcher struct {
	handler AccountFetcherFunc
}

func (t accountFetcher) FetchAccount(ctx context.Context, provider *Provider, account Account) (*Account, error) {
	return t.handler(ctx, provider, account)
}
