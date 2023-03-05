package transferkit

import (
	"sync"
)

type StaticRegistry struct {
	mu        sync.Mutex
	providers map[string]ProviderBuilder
}

func (r *StaticRegistry) Register(code string, p ProviderBuilder) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.providers == nil {
		r.providers = make(map[string]ProviderBuilder)
	}

	r.providers[code] = p
}

func (r *StaticRegistry) Lookup(criteria *RoutingCriteria) (*Provider, error) {
	buildProvider, ok := r.providers[criteria.Code]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return buildProvider(), nil
}

func (r *StaticRegistry) Walk(fn func(code string, builder *Provider) error) error {
	for code, buildProvider := range r.providers {
		if err := fn(code, buildProvider()); err != nil {
			return err
		}
	}

	return nil
}

func (r *StaticRegistry) Entries() []*Provider {
	providers := make([]*Provider, 0)

	for _, buildProvider := range r.providers {
		providers = append(providers, buildProvider())
	}

	return providers
}

type RoutingCriteria struct {
	Code      string
	Countries []string
}
