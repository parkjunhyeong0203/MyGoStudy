package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) //ctx는 만들어진 컨텍스트, cancel은 취소할수있는 펑션. whitCancle()로 만들수 있다.
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
