package main

import (
	"fmt"
	"time"

	prexis "./Interfaces"

	config "./Configuration"
)

func main() {

	config.Init()

	authenticate := prexis.Authenticate()

	time.Sleep(2)

	balance := prexis.CreditBalance(authenticate.Jwt)

	fmt.Println("Balance is: ", balance.Result.Credits)

}
