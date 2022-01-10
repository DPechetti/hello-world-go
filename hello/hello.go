package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Donovan")

	fmt.Println(message)
}
