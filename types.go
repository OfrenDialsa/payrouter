package payrouter

type PaymentMethod string

const (
	PaymentMethodQRIS           PaymentMethod = "qris"
	PaymentMethodVirtualAccount PaymentMethod = "virtual_account"
	PaymentMethodCard           PaymentMethod = "card"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentSucceeded PaymentStatus = "succeeded"
	PaymentFailed    PaymentStatus = "failed"
	PaymentExpired   PaymentStatus = "expired"
	PaymentCanceled  PaymentStatus = "canceled"
)

type CreatePaymentRequest struct {
	OrderID        string
	Amount         int64
	Currency       string
	Method         PaymentMethod
	IdempotencyKey string
}

type Payment struct {
	ID             string
	OrderID        string
	Amount         int64
	Currency       string
	Method         PaymentMethod
	Status         PaymentStatus
	Provider       string
	ProviderID     string
	ProviderStatus string
}
