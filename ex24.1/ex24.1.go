package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hangul := []rune{'가', '나', '다', '라'}
	for _, v := range hangul {
		time.Sleep(300 * time.Millisecond) // 1sec == 1000 millisecond
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 4; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func main() {
	go PrintHangul()
	go PrintNumbers()
	time.Sleep(3 * time.Second)
}
