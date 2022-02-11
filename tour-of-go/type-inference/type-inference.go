package main

import "fmt"

func main() {
	vInt := 4
	vFloat := 42.3
	vComplex := 42.3 + 5i

	fmt.Printf("vInt is of type %T\n", vInt)
	fmt.Printf("vFloat is of type %T\n", vFloat)
	fmt.Printf("vComplex is of type %T\n", vComplex)
}
