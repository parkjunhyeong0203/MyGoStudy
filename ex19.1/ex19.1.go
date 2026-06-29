package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { //함수
	a.balance -= amount
}
func (a *account) withdrawMethod(amount int) { //메서드
	a.balance -= amount
}

func main() {
	a := &account{100} //포인터 타입 으로 선언
	withdrawFunc(a, 30)
	a.withdrawMethod(30)

	fmt.Printf("%d\n", a.balance)
}
