package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합은 %d\n", a, b, sum)
	wg.Done()
}
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000)
	}
	wg.Wait()

	fmt.Println("모든 작업 완료")
}
