package payrouter

import "errors"

var (
	ErrProviderUnavailable       = errors.New("provider unavailable")
	ErrPaymentRejected           = errors.New("payment rejected")
	ErrPaymentNotFound           = errors.New("payment not found")
	ErrProviderPaymentNotSupport = errors.New("no provider supports this payment method")
)
