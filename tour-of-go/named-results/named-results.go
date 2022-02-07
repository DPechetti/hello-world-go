package main

import "fmt"

func split(value int) (x, y int) {
	x = value * 2 / 3
	y = value - x
	return
}

func main() {
	fmt.Println(split(9))
}
