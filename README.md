# angel-stocks

Go library for fetching historical data from AngelOne SmartAPI

> This is not the official library. I made this when I was learning GoLang.

## Usage

```go
package main

import (
	"fmt"

	stocks "github.com/aryan828/angel-stocks"
	"github.com/xlzd/gotp"
)

func main() {
	stock := stocks.NewClient(
        "ANGEL_CLIENT_ID",
        "mPIN",
        gotp.NewDefaultTOTP("TOTP_STRING").Now(),
        "API_KEY",
        )
	_, err := stock.Connect()
	if err != nil {
		return
	}

	profile, err := stock.GetProfile()
	if err != nil {
		return
	}
	fmt.Println(profile)

	history, err := stock.History(
        "EXCHANGE",       // stocks.NSE, stocks.NFO
        "SYMBOL",         
        "INTERVAL",       // stocks.ONE_MINUTE, stocks.THREE_MINUTE, ...
        "FROM_DATE_TIME", // "2021-06-10 09:15"
        "TO_DATE_TIME",
        )
	if err != nil {
		return
	}
	fmt.Println(history)
}

```
