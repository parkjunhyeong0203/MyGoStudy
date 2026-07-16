package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var fibMap [65535]int

/*func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}*/

func Fib(n int) int {
	f := fibMap[n]
	if f > 0 {
		return f
	}
	if n == 0 {
		f = 0
	} else if n == 1 {
		f = 1
	} else {
		f = Fib(n-1) + Fib(n-2)
	}
	fibMap[n] = f
	return f
}

func main() {
	f, err := os.Create("cpu1.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	fmt.Println(Fib(50))

	time.Sleep(10 * time.Second)
}
