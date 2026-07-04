package main

import "fmt"

func main() {
	m := make(map[string]string) //initialize, go는 HashMap기반

	m["이화랑"] = "서울시 광진구" //순서보장 X
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구"

	fmt.Printf("송하나의 주소는 %s 입니다\n", m["송하나"])
	fmt.Printf("백두산의 주소는 %s 입니다\n", m["백두산"])

	for k, v := range m {
		fmt.Println(k, v)
	}

	if _, ok := m["백두산"]; ok {
		delete(m, "백두산")      //삭제하면 기본값으로 초기화, 그래서 ok로 확인해야된다
		fmt.Println(m["백두산"]) //nil로 초기화 된듯?
	}
}
