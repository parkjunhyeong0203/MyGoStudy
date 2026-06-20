package main

import "fmt"

func main() {
	str := "Hello 월드"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	}
	for _, v := range str { //이렇게도 가능, 다만 for문에서만
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", v, v, v)

	}
}
