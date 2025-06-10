package main

import (
	"fmt"

	"github.com/ksundev/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ksun")
	fmt.Println(account)
}
