package provider

import (
	"context"

	"github.com/OfrenDialsa/payrouter"
)

type Provider interface {
	Name() string
	CreatePayment(ctx context.Context, req payrouter.CreatePaymentRequest) (*payrouter.Payment, error)
	GetPayment(ctx context.Context, id string) (*payrouter.Payment, error)
}
