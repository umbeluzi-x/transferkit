package transferkit

import (
	"sync"
)

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{
		providers: make(map[string]ProviderBuilder),
	}
}

type ProviderRegistry struct {
	mu        sync.Mutex
	providers map[string]ProviderBuilder
}

func (r *ProviderRegistry) Register(code string, p ProviderBuilder) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.providers == nil {
		r.providers = make(map[string]ProviderBuilder)
	}

	r.providers[code] = p
}

func (r *ProviderRegistry) Lookup(code string) (*Provider, error) {
	if r.providers == nil {
		return nil, ErrProviderNotFound
	}

	buildProvider, ok := r.providers[code]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return buildProvider(), nil
}

func (r *ProviderRegistry) Walk(walkFn func(code string, provider *Provider) error) error {
	for code, buildProvider := range r.providers {
		if err := walkFn(code, buildProvider()); err != nil {
			return err
		}
	}

	return nil
}

func (r *ProviderRegistry) Providers() []*Provider {
	providers := make([]*Provider, 0)

	for _, buildProvider := range r.providers {
		providers = append(providers, buildProvider())
	}

	return providers
}
