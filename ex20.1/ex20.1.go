package main

import "fmt"

type Stringer interface { // String() 메서드가 있다면 Stringer 타입으로 보겠다.
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("안녕, 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer
	stringer = &student //implement? 객체가 인터페이스의 메서드를 포함하고 있으면 가르킬수 있다. String()만 호출가능

	fmt.Printf("%s\n", stringer.String())
}
