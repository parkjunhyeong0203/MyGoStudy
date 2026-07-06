package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) { //scanner 객체는 데이터를 한줄씩 읽어오기 편하다.
	scanner := bufio.NewScanner(strings.NewReader(str)) //NewScanner는 io.Reader 타입(Read함수를 구현한 것)을 인자로 받는다.
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner) //첫번째 단어 읽기
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err) //%w는 에러를 래핑
	}
	pos += n + 1
	b, n, err := readNextInt(scanner) // 두번째 단어 읽기
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err) //새로운 에러 객체 만듬, 기존 에러 %w는 에러를 래핑
	}
	return a * b, nil
}
func readNextInt(scanner *bufio.Scanner) (int, int, error) { //단어를 읽어서 숫자로 변환후 반환, 변환된 숫자, 읽은 글자수, 에러
	if !scanner.Scan() { //단어읽기 실패했을때
		return 0, 0, fmt.Errorf("failed to scan")
	}
	word := scanner.Text()            //읽어온걸(scan 한거를) 문자열로
	number, err := strconv.Atoi(word) // "24" -> 24, "abc" -> error, NumError 반환함
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s, err:%w, ", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { // wrapping 한거 unboxing, 변환이 가능하다면 에러가 났다는것
			fmt.Println("NUmberError", numError)
		}
	}
}
func main() {
	readEq("123 3")
	readEq("123 abc")
	//panic(...) 모든 타입 가능, panic은 복구를 안 할 경우 강제 종료, 보통 개발시 많이 사용. 다른 에러 핸들링은 배포후 사용
	//recover(...) 복구 함수, 패닉 객체를 반환, 왠만하면 안하는게 좋다
	/* defer func() {
		if r := recover(); r != nil {
			//fmt.Println("panic recover - ", r)
		}
	   }() */ //()이건 함수 리터럴(람다식)호출하는것
}
