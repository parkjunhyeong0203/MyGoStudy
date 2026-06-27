package main

import "fmt"

func addNum(slice []int) {
	slice = append(slice, 4)
}

func main() {
	slice := []int{1, 2, 3}

	addNum(slice)
	fmt.Println(slice) // 그대로 1 2 3 출력
	// 함수 호출시 slice = slice 로 복사. slice 내부 배열을 가르키는 포인터가 같은 주소를 가르킴. 인스턴스는 다르다.
	// 그런데 공간이 없으므로 append 할때 다시 slice가 새롭게 만들어진다.
	// 공간이 있어도 addNum의 slice에만 길이가 4고 main을 여전히 3을 가르킨다.
} // 따라서 append를 해도 기존 slice는 변화가 없음. 포인터 or return 쓰면 바뀐다.
