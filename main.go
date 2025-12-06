package main

import (
	"fmt"

	"example.com/banking/accounts"
)

func main() {
	account := accounts.NewAccount("gyusun")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(15)
	if err != nil {
		fmt.Println("error 내용 ", err)
	}
	fmt.Println(account.Balance())

}
