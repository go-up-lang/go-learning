package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 입금
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// 출금
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a Account) Balance() int {
	return a.balance
}

// new struct
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	// return "whatever you want"
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
