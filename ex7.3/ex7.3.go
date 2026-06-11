package main

import (
	"fmt"
)

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return

}
func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0) //이미 있는 변수 또 선언? 여러개 선언할때 선언 대입문 쓸때는 하나라도 새로운 변수 있을때 쓰게해줌
	fmt.Println(d, success)
}
