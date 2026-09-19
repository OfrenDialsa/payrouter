package router

import (
	"github.com/OfrenDialsa/payrouter"
	"github.com/OfrenDialsa/payrouter/provider"
)

type Strategy interface {
	Select(providers []provider.Provider, req payrouter.CreatePaymentRequest) (provider.Provider, error)
}
