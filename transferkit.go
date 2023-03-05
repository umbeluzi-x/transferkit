package transferkit

import (
	"context"
	"errors"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"sync"
)

var (
	ErrInvalidDeliveryMethodType = errors.New("invalid payment method")
	ErrInvalidCurrency           = errors.New("invalid currency")
	ErrConnectionTimeout         = errors.New("connection timeout")
	ErrConnectionRefused         = errors.New("connection refused")
	ErrInvalidProviderCode       = errors.New("invalid provider code")
	ErrProviderNotFound          = errors.New("provider not found")
	MissingConfigurationHook     = errors.New("missing initializer")
)

const (
	defaultVersion = "latest"
)

type Set struct {
	mu      sync.RWMutex
	entries map[string]struct{}
}

func NewSet(s ...string) *Set {
	set := &Set{
		entries: make(map[string]struct{}),
	}

	set.mu.Lock()
	defer set.mu.Unlock()

	for _, i := range s {
		set.entries[i] = struct{}{}
	}

	return set
}

func (s *Set) Add(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.entries[v] = struct{}{}
}

func (s *Set) Remove(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.entries[v]; ok {
		delete(s.entries, v)
	}
}

type ProviderBuilder func() *Provider

type Provider struct {
	Code               string
	Icon               string
	Description        string
	Author             string
	Version            string
	RequireConfig      string
	RequireSecrets     *Set
	DeliveryMethod     *Set
	OnConfigure        ProviderConfigurator
	OnTransactionSend  TransactionSender
	OnTransactionCheck TransactionChecker
	OnAccountCheck     AccountChecker
	OnAccountFetch     AccountFetcher
	OnBalanceFetch     BalanceFetcher
	OnClose            ProviderCloser
}

type Secret map[string]string

func (p *Provider) Configure(ctx context.Context, opts *ProviderOptions) error {
	if p.OnConfigure == nil {
		return MissingConfigurationHook
	}

	return p.OnConfigure.Configure(ctx, opts)
}

func (p *Provider) Close(ctx context.Context) error {
	if p.OnClose == nil {
		return nil
	}

	return p.OnClose.Close(ctx)
}

func (p Provider) Validate() error {
	if p.Code == "" {
		return ErrInvalidProviderCode
	}

	return nil
}

func (p Provider) String() string {
	x := `# {{.Name}}
## Implements
{{ range v := .Capabilities }}
* {{ $v }}
{{ end }}

## Config
{{ .RequireConfig }}`

	return x
}

type Capability string

const (
	CapSendTransaction  = Capability("CAP_SEND_TRANSACTION")
	CapCheckTransaction = Capability("CAP_CHECK_TRANSACTION")
	CapCheckAccount     = Capability("CAP_CHECK_ACCOUNT")
	CapFetchAccount     = Capability("CAP_FETCH_ACCOUNT")
	CapFetchBalance     = Capability("CAP_FECTH_BALANCE")
)

func (p Provider) Capabilities() []Capability {
	capabilities := make([]Capability, 0)

	if p.OnTransactionSend != nil {
		capabilities = append(capabilities, CapSendTransaction)
	}

	if p.OnTransactionCheck != nil {
		capabilities = append(capabilities, CapCheckTransaction)
	}

	if p.OnAccountCheck != nil {
		capabilities = append(capabilities, CapCheckAccount)
	}

	if p.OnAccountFetch != nil {
		capabilities = append(capabilities, CapFetchAccount)
	}

	if p.OnBalanceFetch != nil {
		capabilities = append(capabilities, CapFetchBalance)
	}

	return capabilities
}

type ProviderOptions struct {
	Config  []byte
	Cache   Cache
	Secrets Secret
}

type Cache interface{}

type ProviderConfiguratorFunc func(context.Context, *ProviderOptions) error

func (c ProviderConfiguratorFunc) Configure(ctx context.Context, opts *ProviderOptions) error {
	return c(ctx, opts)
}

type ProviderConfigurator interface {
	Configure(context.Context, *ProviderOptions) error
}

type CloserFunc func(context.Context) error

func (c CloserFunc) Close(ctx context.Context) error {
	return c(ctx)
}

type ProviderCloser interface {
	Close(context.Context) error
}

func NewSchema(v string) *Schema {
	newSchema := &Schema{
		value: jsonschema.MustCompileString("config.json", v),
	}

	return newSchema
}

type Schema struct {
	value *jsonschema.Schema
}
