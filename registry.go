package transferkit

import (
	"errors"
	"sync"
)

type Registry struct {
	mu        sync.Mutex
	providers map[string]ProviderBuilder
}

func (r *Registry) Register(code string, p ProviderBuilder) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.providers == nil {
		r.providers = make(map[string]ProviderBuilder)
	}

	r.providers[code] = p
}

func (r *Registry) Find(code string) (ProviderBuilder, error) {
	builder, ok := r.providers[code]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return builder, nil
}

func (r *Registry) Walk(fn func(code string, builder ProviderBuilder) error) error {
	for code, builder := range r.providers {
		if err := fn(code, builder); err != nil {
			return err
		}
	}

	return nil
}
