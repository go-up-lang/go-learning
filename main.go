package main

import (
	"fmt"

	"example.com/banking/accounts"
)

func main() {
	account := accounts.NewAccount("gyusun")
	fmt.Println(account)
}
