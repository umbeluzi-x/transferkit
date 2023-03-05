package transferkit

import (
	"context"
	"errors"
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

type ProviderBuilder func() *Provider

type Provider struct {
	Code               string
	Description        string
	Author             string
	Version            string
	Config             []byte
	OnConfigure        ProviderConfigurator
	OnTransactionSend  TransactionSender
	OnTransactionCheck TransactionChecker
	OnAccountCheck     AccountChecker
	OnAccountFetch     AccountFetcher
	OnBalanceFetch     BalanceFetcher
	OnClose            ProviderCloser
}

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

type ProviderOptions struct {
	Config []byte
	Cache  Cache
}

type Cache interface{}

type ConfiguratorFunc func(context.Context, *ProviderOptions) error

func (c ConfiguratorFunc) Configure(ctx context.Context, opts *ProviderOptions) error {
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
