package transferkit

import "context"

type AccountCheckerFunc func(ctx context.Context, provider *Provider, transaction AccountCheck) (*AccountCheckResponse, error)

type AccountChecker interface {
	CheckAccount(ctx context.Context, provider *Provider, transaction AccountCheck) (*AccountCheckResponse, error)
}

func (a AccountCheckerFunc) CheckAccount(ctx context.Context, provider *Provider, account AccountCheck) (*AccountCheckResponse, error) {
	return a(ctx, provider, account)
}

type AccountCheckResponse struct{}
