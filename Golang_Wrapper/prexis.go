package main

import (
	"fmt"

	config "./Configuration"
	prexis "./Interfaces"
)

func main() {

	config.Init()

	authenticate := prexis.Authenticate()

	balance := prexis.CreditBalance(authenticate.Jwt)
	hash := prexis.RegisterHash(authenticate.Jwt, "asasdasdasdad")

	fmt.Println("Balance is: ", balance.Result.Credits)
	fmt.Println("Result is: ", hash.Result.Message)

}
