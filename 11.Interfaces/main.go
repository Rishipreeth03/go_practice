package main

import (
	"fmt"
)

// Define an interface that represents any payment gateway
// Following the Open/Closed Principle — code should be open for extension but closed for modification
type paymenter interface {
	pay(amount float32)
}

// The payment struct depends on the 'paymenter' interface (abstraction),
// not on a specific implementation like Razorpay.
type payment struct {
	gateway paymenter
}

// makePayment calls the 'pay' method of whichever payment gateway is assigned.
// This allows new gateways (like PayPal, Stripe, etc.) to be added
// without changing existing code.
func (p payment) makePayment(amount float32) {
	// Previously, Razorpay was hardcoded like this:
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// Now we delegate the payment logic to the interface implementation.
	p.gateway.pay(amount)
}

// Concrete type implementing the 'paymenter' interface
type razorpay struct{}

// 'pay' method implementation for Razorpay
func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using Razorpay:", amount)
}

// You could easily add another payment gateway without changing the main logic:
// type paypal struct{}
// func (p paypal) pay(amount float32) {
//     fmt.Println("Making payment using PayPal:", amount)
// }

func main() {
	// Create a Razorpay gateway instance
	razorpayPaymentGw := razorpay{}

	// Create a Payment object that uses Razorpay as its payment gateway
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}

	// Make the payment using the chosen gateway
	newPayment.makePayment(100)
}
