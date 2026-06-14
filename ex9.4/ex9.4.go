package main

import "fmt"

func getMyAge() (int, bool) { //private 선언
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("you are young")
	} else if ok && age < 30 {
		fmt.Println("good age")
	} else if ok {
		fmt.Println("old man")
	} else {
		fmt.Println("error")
	}
}
