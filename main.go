package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := []string{"nico", "gyusun"}
	for _, person := range people {
		go isSexy(person, c)
	}

	result1 := <-c
	result2 := <-c
	fmt.Println(result1)
	fmt.Println(result2)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	if person == "nico" {
		c <- person + " false"
	} else {
		c <- person + " true"
	}
}
