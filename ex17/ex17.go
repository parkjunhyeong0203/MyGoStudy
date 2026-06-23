package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	cnt := 0
	r := rand.Intn(100)
	for {
		fmt.Printf("정수 숫자를 입력하세요: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("제대로 입력되지 않았습니다.")
		} else {
			if n == r {
				fmt.Println("숫자를 맞추셨습니다!")
				break
			} else if n > r {
				fmt.Println("입력한 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력한 숫자가 더 작습니다.")
			}
		}
		cnt++
	}
	fmt.Println("시도한 횟수:", cnt)

}
