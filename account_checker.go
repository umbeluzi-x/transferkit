package transferkit

import "context"

type AccountCheckerFunc func(ctx context.Context, provider *Provider, transaction Account) (*Account, error)

type AccountChecker interface {
	CheckAccount(ctx context.Context, provider *Provider, transaction Account) (*Account, error)
}

func (a AccountCheckerFunc) CheckAccount(ctx context.Context, provider *Provider, account Account) (*Account, error) {
	return a(ctx, provider, account)
}
