package main

import (
	"fmt"

	"example.com/banking/banking"
)

func main() {
	account := banking.Account{Owner: "gyusun", Balance: 3}
	fmt.Println(account)
}
