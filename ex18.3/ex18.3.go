package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2] //1번째 인덱스 부터 1번쨰 인덱스 까지, 기존 배열에서 가져온다.
	fmt.Println(array)
	fmt.Println(slice, len(slice), cap(slice))

	array[1] = 100
	fmt.Println("after change value: ", slice) //기존 배열 or 슬라이스를 가르키니 같이 바뀐다. 그냥 가져와서 len이랑 cap만 설정하는듯

	slice = append(slice, 500)
	fmt.Println("after append 500 slice:", slice, "and array:", array) //공간이 차있으면 배열은 안바뀔듯?
}
