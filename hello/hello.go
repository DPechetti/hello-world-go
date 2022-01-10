package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Donovan")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
