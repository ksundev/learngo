package accounts

import (
	"errors"
	"fmt"
)

// Account represents a bank account
type Account struct {
	// private members
	owner   string
	balance int
}

// NewAccount creates a new account
func NewAccount(owner string) *Account {
	return &Account{
		owner:   owner,
		balance: 0,
	}
}

// Deposit adds amount to the balance
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw from the balance
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("can't withdraw, insufficient balance")
	}
	a.balance -= amount
	return nil
}

func (a Account) String() string {
	return fmt.Sprintf("Account owner: %s, Balance: %d", a.owner, a.balance)
}
