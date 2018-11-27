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
	hash := prexis.RegisterHash(authenticate.Jwt, "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387")
	transaction := prexis.GetTransaction(authenticate.Jwt, "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387")

	fmt.Println("Balance is: ", balance.Result.Credits)
	fmt.Println("Result is: ", hash.Result.Message)
	fmt.Println("Transaction is: ", transaction.Result.Message)

}
