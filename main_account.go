package main

import (
	"fmt"
	"log"

	"github.com/ksundev/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ksun")
	account.Deposit(40)
	fmt.Println(account)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)

	account.ChangeOwner("kikko")
	fmt.Println(account.GetOwner())
}
