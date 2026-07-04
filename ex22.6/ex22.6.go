package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]string{}

	m[hash(23)] = "송하나"
	m[hash(259)] = "백두산"
	fmt.Printf("%d = %s\n", 23, m[hash(23)])
	fmt.Printf("%d = %s\n", 259, m[hash(259)])

}
