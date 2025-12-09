package main

import (
	"fmt"

	"example.com/banking/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	definitionValue, _ := dictionary.Search(word)
	fmt.Println(definitionValue)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(definition)

}
