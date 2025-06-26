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

	// Update Test
	baseword := "first"
	newDef := "New first word"
	wrongword := "wrong"
	// Try to update a wrong word
	err = dictionary.Update(wrongword, newDef)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("This should not happen")
	}

	// Try to update a correct word
	err = dictionary.Update(baseword, newDef)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(baseword, "word updated")
		definition, err = dictionary.Search(baseword)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	}

	// Delete Test
	// Try to delete a wrong word
	wrongword = "wrong"
	err = dictionary.Delete(wrongword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("This should not happen")
	}

	// Try to delete a correct word
	word = "second"
	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word, "word deleted")
	}

	// Try to search a deleted word
	definition, err = dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
