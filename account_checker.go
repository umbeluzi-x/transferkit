package providerkit

import "context"

type AccountCheckerFunc func(ctx context.Context, provider *Provider, payment Account) (*Account, error)

type AccountChecker interface {
	CheckAccount(ctx context.Context, provider *Provider, payment Account) (*Account, error)
}

func AccountCheckFunc(f AccountCheckerFunc) AccountChecker {
	return accountChecker{
		handler: f,
	}
}

type accountChecker struct {
	handler AccountCheckerFunc
}

func (t accountChecker) CheckAccount(ctx context.Context, provider *Provider, transaction Account) (*Account, error) {
	return t.handler(ctx, provider, transaction)
}
