package main

import "fmt"

func main() {
	sum := 1

	for sum < 1000 { //lint removed semicolon
		sum += sum
	}

	fmt.Println(sum)
}
