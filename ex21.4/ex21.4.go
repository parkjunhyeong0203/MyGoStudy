package main

import "fmt"

type opFn func(int, int) int

func getOperator(op string) opFn {
	if op == "+" {
		return func(a, b int) int { //람다식
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4)
	fmt.Println(result)
}
