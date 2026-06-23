package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. 키보드 입력용 버퍼 리더 생성
	stdin := bufio.NewReader(os.Stdin)

	fmt.Print("이름을 입력하세요 (공백 포함 가능): ")

	// 2. 엔터키('\n')가 나올 때까지 한 줄을 통째로 읽어옴
	line, _ := stdin.ReadString('\n')

	fmt.Printf("입력하신 이름은: %s", line)
}
