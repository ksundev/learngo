package accounts

import "fmt"

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

func (a Account) String() string {
	return fmt.Sprintf("Account owner: %s, Balance: %d", a.owner, a.balance)
}
