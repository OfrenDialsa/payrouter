package router

import (
	"github.com/OfrenDialsa/payrouter"
	"github.com/OfrenDialsa/payrouter/provider"
)

type Router struct {
	providers map[string]provider.Provider
	strategy  Strategy
}

func New(strategy Strategy) *Router {
	return &Router{
		providers: make(map[string]provider.Provider),
		strategy:  strategy,
	}
}

func (r *Router) Register(p provider.Provider) {
	r.providers[p.Name()] = p
}

func (r *Router) Route(req payrouter.CreatePaymentRequest) (provider.Provider, error) {
	providers := make([]provider.Provider, 0, len(r.providers))

	for _, p := range r.providers {
		providers = append(providers, p)
	}

	selected, err := r.strategy.Select(providers, req)
	if err != nil {
		return nil, err
	}

	if selected == nil {
		return nil, payrouter.ErrProviderUnavailable
	}

	return selected, nil
}
