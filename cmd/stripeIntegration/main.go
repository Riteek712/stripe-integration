package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/customer"
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
}
