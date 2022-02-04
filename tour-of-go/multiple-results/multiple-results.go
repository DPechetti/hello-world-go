package main

import "fmt"

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("Donovan", "Pechetti")
	fmt.Println(a, b)
}
