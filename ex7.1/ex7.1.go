package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	c := add(3, 6)
	fmt.Println(c)
}
