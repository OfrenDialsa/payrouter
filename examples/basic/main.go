package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/OfrenDialsa/payrouter"
// )

// func main() {
// 	router := payrouter.New()

// 	payment, err := router.CreatePayment(
// 		context.Background(),
// 		payrouter.CreatePaymentRequest{
// 			OrderID:  "ORDER-001",
// 			Amount:   100000,
// 			Currency: "IDR",
// 			Method:   payrouter.PaymentMethodQRIS,
// 		},
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(payment)
// }
