package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	slice3 := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	//slice2 := append([]int{}, slice1...) 이렇게 하면 순회 안하고 복사

	copy(slice2, slice1) //내장 함수 써서 복사, 첫번쨰 인자가 목적지, 두번쨰인자가 소스

	/*for i, v := range slice1 { //순회를 이용해서 복사
		slice2[i] = v
	} */

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice3 = append(slice3[:idx], slice3[idx+1:]...)                      //원하는 idx 삭제하기
	slice3 = append(slice3[:idx], append([]int{100}, slice3[idx:]...)...) //원하는 idx에 삽입
	fmt.Println(slice3)
}
