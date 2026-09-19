package midtrans

import (
	"context"

	"github.com/OfrenDialsa/payrouter"
	"github.com/OfrenDialsa/payrouter/provider"
)

type Provider struct {
	serverKey string
}

func New(serverKey string) *Provider {
	return &Provider{
		serverKey: serverKey,
	}
}

func (p *Provider) Name() string {
	return "midtrans"
}

func (p *Provider) Capabilities() provider.Capabilities {
	return provider.Capabilities{
		Methods: []payrouter.PaymentMethod{
			payrouter.PaymentMethodCard,
			payrouter.PaymentMethodQRIS,
			payrouter.PaymentMethodVirtualAccount,
		},
	}
}

func (p *Provider) CreatePayment(ctx context.Context, req payrouter.CreatePaymentRequest) (*payrouter.Payment, error) {
	// TODO: call Midtrans API
	return nil, nil
}

func (p *Provider) GetPayment(ctx context.Context, id string) (*payrouter.Payment, error) {
	// TODO: call Midtrans API status
	return nil, nil
}
