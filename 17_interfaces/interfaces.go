package main

import "fmt"

// interfaces are like a contract (add er)
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	// gateway razorpay // makes testing difficult since this is hard coded --> solution:interfaces
	gateway paymenter
}

// Open close principe violated -->open for extension closed for modification
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// the already written code we had to modify --> violates solid principle --> open close principle
	// stripePaymentGw := stripe{}
	p.gateway.pay(amount)
}

type razorpay struct{}

// In Go , we dont have implements, if the method defined in the interface has the same signature here
// it thinks that it is implemnting that method here ...
func (r razorpay) pay(amount float32) {
	//logic to make payment using razorpay apis/sdk
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making paymetn using fake gateway for testing purposes")
}

func (f fakepayment) refund(amount float32, account string) {
	fmt.Println("making refund through the fake gateway for testing purposes")
}
func main() {
	fmt.Println("Interfaces")
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	fakeGw := fakepayment{}
	newPayment := payment{
		gateway: fakeGw,
	}
	newPayment.makePayment(100)
}
