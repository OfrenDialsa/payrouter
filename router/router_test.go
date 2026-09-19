package router

import (
	"context"
	"testing"

	"github.com/OfrenDialsa/payrouter"
	"github.com/OfrenDialsa/payrouter/provider"
)

type MockProvider struct {
	name         string
	capabilities provider.Capabilities
}

func (m *MockProvider) Name() string {
	return m.name
}

func (m *MockProvider) Capabilities() provider.Capabilities {
	return m.capabilities
}

func (m *MockProvider) CreatePayment(ctx context.Context, req payrouter.CreatePaymentRequest) (*payrouter.Payment, error) {
	return nil, nil
}

func (m *MockProvider) GetPayment(ctx context.Context, id string) (*payrouter.Payment, error) {
	return nil, nil
}

func TestPriorityStrategy_Select(t *testing.T) {
	tests := []struct {
		name             string
		priority         []string
		providers        []provider.Provider
		requestMethod    payrouter.PaymentMethod
		expectedProvider string
		expectError      bool
	}{
		{
			name:     "select first provider by priority",
			priority: []string{"doku", "midtrans"},
			providers: []provider.Provider{
				&MockProvider{
					name: "doku",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodQRIS,
						},
					},
				},
				&MockProvider{
					name: "midtrans",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodQRIS,
						},
					},
				},
			},
			requestMethod:    payrouter.PaymentMethodQRIS,
			expectedProvider: "doku",
		},

		{
			name:     "skip provider that does not support method",
			priority: []string{"doku", "midtrans"},
			providers: []provider.Provider{
				&MockProvider{
					name: "doku",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodVirtualAccount,
						},
					},
				},
				&MockProvider{
					name: "midtrans",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodQRIS,
						},
					},
				},
			},
			requestMethod:    payrouter.PaymentMethodQRIS,
			expectedProvider: "midtrans",
		},

		{
			name:     "return error when no provider supports method",
			priority: []string{"doku", "midtrans"},
			providers: []provider.Provider{
				&MockProvider{
					name: "doku",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodVirtualAccount,
						},
					},
				},
				&MockProvider{
					name: "midtrans",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodCard,
						},
					},
				},
			},
			requestMethod: payrouter.PaymentMethodQRIS,
			expectError:   true,
		},

		{
			name:     "return error when priority provider is not registered",
			priority: []string{"doku", "midtrans"},
			providers: []provider.Provider{
				&MockProvider{
					name: "xendit",
					capabilities: provider.Capabilities{
						Methods: []payrouter.PaymentMethod{
							payrouter.PaymentMethodQRIS,
						},
					},
				},
			},
			requestMethod: payrouter.PaymentMethodQRIS,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := NewPriorityStrategy(tt.priority...)

			req := payrouter.CreatePaymentRequest{
				Method: tt.requestMethod,
			}

			selected, err := strategy.Select(
				tt.providers,
				req,
			)

			if tt.expectError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if selected == nil {
				t.Fatal("expected provider, got nil")
			}

			if selected.Name() != tt.expectedProvider {
				t.Fatalf(
					"expected provider %q, got %q",
					tt.expectedProvider,
					selected.Name(),
				)
			}
		})
	}
}

func TestRouter_Register(t *testing.T) {
	strategy := NewPriorityStrategy("doku")

	r := New(strategy)

	doku := &MockProvider{
		name: "doku",
		capabilities: provider.Capabilities{
			Methods: []payrouter.PaymentMethod{
				payrouter.PaymentMethodQRIS,
			},
		},
	}

	r.Register(doku)

	selected, err := r.Route(
		payrouter.CreatePaymentRequest{
			Method: payrouter.PaymentMethodQRIS,
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if selected == nil {
		t.Fatal("expected provider, got nil")
	}

	if selected.Name() != "doku" {
		t.Fatalf(
			"expected doku, got %q",
			selected.Name(),
		)
	}
}

func TestRouter_NoProviderRegistered(t *testing.T) {
	strategy := NewPriorityStrategy("doku")

	r := New(strategy)

	_, err := r.Route(
		payrouter.CreatePaymentRequest{
			Method: payrouter.PaymentMethodQRIS,
		},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
