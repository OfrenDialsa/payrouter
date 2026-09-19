package router

import (
	"slices"

	"github.com/OfrenDialsa/payrouter"
	"github.com/OfrenDialsa/payrouter/provider"
)

type PriorityStrategy struct {
	priority []string
}

func NewPriorityStrategy(priority ...string) *PriorityStrategy {
	return &PriorityStrategy{
		priority: priority,
	}
}

func (s *PriorityStrategy) Select(providers []provider.Provider, req payrouter.CreatePaymentRequest) (provider.Provider, error) {
	for _, name := range s.priority {
		for _, p := range providers {
			if p.Name() != name {
				continue
			}

			if !supportsMethod(p, req.Method) {
				continue
			}

			return p, nil
		}
	}

	return nil, payrouter.ErrProviderPaymentNotSupport
}

func supportsMethod(p provider.Provider, method payrouter.PaymentMethod) bool {
	return slices.Contains(p.Capabilities().Methods, method)
}
