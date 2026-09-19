package doku

import (
	"context"

	"github.com/OfrenDialsa/payrouter"
	"github.com/OfrenDialsa/payrouter/provider"
)

type Provider struct {
	clientID string
	secret   string
}

func New(clientID, secret string) *Provider {
	return &Provider{
		clientID: clientID,
		secret:   secret,
	}
}

func (p *Provider) Name() string {
	return "doku"
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
	// TODO: call DOKU API
	return nil, nil
}

func (p *Provider) GetPayment(ctx context.Context, id string) (*payrouter.Payment, error) {
	// TODO: call DOKU API status
	return nil, nil
}
