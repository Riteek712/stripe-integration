package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	// "github.com/stripe/stripe-go/paymentintent"
	// "github.com/stripe/stripe-go/invoice"
)

func main() {
	godotenv.Load()
	secretKey := os.Getenv("SecretKey")
	stripe.Key = secretKey
	fmt.Println("Making Requests")
	// fmt.Println(secretKey)

}
