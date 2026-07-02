package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) { //외부에서 로직 주입 = 의존성 주입
		fmt.Fprintln(f, msg)
	})
	//	  writeHello(func(msg string) {
	//		fmt.Println(msg)
	//	})
}
