package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number) //_은 빈칸지시자. 안쓰고 싶을때 밑줄
		if err != nil {
			fmt.Println("숫자를 입력하세요")

			stdin.ReadString('\n')
			continue
		}
		fmt.Println("입력하센 숫자는 ", number, "입니다")
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문 종료")
}
