package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/customer"
	"github.com/stripe/stripe-go/v75/paymentintent"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatalf("SECRET_KEY not set in .env file")
	}

	// Set the Stripe API key
	stripe.Key = secretKey
	fmt.Println("Making Requests")

	// Create a customer with parameters
	params := &stripe.CustomerParams{
		Name:  stripe.String("Jenny Rosen"),
		Email: stripe.String("jennyrosen@example.com"),
	}
	result, err := customer.New(params)
	if err != nil {
		fmt.Printf("Error creating customer: %v\n", err)
		return
	}
	fmt.Printf("Customer created: %v\n", result)
	fmt.Printf("Customer name: %v\n", result.Name)
	fmt.Printf("Customer email: %v\n", result.Email)

	piParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	piresult, err := paymentintent.New(piParams)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(piresult)

}
