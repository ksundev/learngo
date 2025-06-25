package main

import (
	"fmt"

	"github.com/ksundev/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "안녕"
	dictionary["world"] = "세상"
	dictionary["go"] = "고"

	fmt.Println(dictionary["hello"])
}
