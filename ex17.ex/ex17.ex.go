package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	money := 1000
	for {
		fmt.Printf("1 ~ 5 정수 숫자를 입력하세요: ")
		n, err := InputIntValue()
		r := rand.Intn(5) + 1
		if err != nil {
			fmt.Println("제대로 입력되지 않았습니다.")
		} else {
			if n == r {
				money = money + 500
				fmt.Println("축하합니다! 잔액:", money)
			} else {
				money = money - 100
				fmt.Println("아쉽습니다.. 잔액:", money)
			}
		}
		if money <= 0 || money >= 5000 {
			break
		}
	}

}
