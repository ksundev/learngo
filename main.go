package main

import (
	"fmt"

	"github.com/ksundev/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	word := "second"
	def := "Second word"
	err = dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word, "word added")
	}
}
