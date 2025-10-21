package main

import "fmt"

// OldPaymentProcessor legacy payment
type OldPaymentProcessor struct {
}

func (o *OldPaymentProcessor) Process(sum float64) {
	fmt.Printf("Payment amount %v processed by OldPaymentProcessor\n", sum)
}

// NewPaymentProcessor actual payment mechanism
type NewPaymentProcessor interface {
	Pay(amount float64)
}

// PaymentAdapter adapter
type PaymentAdapter struct {
	oldPaymentProcessor *OldPaymentProcessor
}

func (p *PaymentAdapter) Pay(amount float64) {
	p.oldPaymentProcessor.Process(amount)
}

func main() {
	payment := &PaymentAdapter{oldPaymentProcessor: &OldPaymentProcessor{}}
	payment.Pay(999.99)
}
