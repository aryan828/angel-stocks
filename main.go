package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/xlzd/gotp"
)

func main() {
	err := godotenv.Load("cred.env")
	if err != nil {
		return
	}

	var stock Client
	stock.credentials = map[string]string{
		"clientcode": os.Getenv("CLIENT_CODE"),
		"password":   os.Getenv("PASSWORD"),
		"totp":       gotp.NewDefaultTOTP(os.Getenv("TOTP")).Now(),
	}
	stock.apiKey = os.Getenv("API_KEY")

	tokens, err := stock.Connect()
	if err != nil {
		return
	}
	fmt.Println(tokens)

	profile, err := stock.GetProfile()
	if err != nil {
		return
	}
	fmt.Println(profile)
}
