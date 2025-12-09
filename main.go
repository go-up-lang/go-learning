package main

import (
	"fmt"

	"example.com/banking/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First1")

	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 		fmt.Println(err)
	// 	}

	err := dictionary.Delete(baseWord)
	word, _ := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)

}
