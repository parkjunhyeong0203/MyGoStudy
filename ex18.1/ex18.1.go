package main

import "fmt"

func main() {
	//slice := make([]int, 3) 슬라이스(동적 배열) 만드는법
	slice := []int{1, 2, 3}
	slice1 := append(slice, 4)        //밸류 추가하고 새로운 슬라이스 반환
	for i := 0; i < len(slice); i++ { //순회법
		slice[i] += 10
	}
	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println(slice)
	fmt.Println(slice1)
}
