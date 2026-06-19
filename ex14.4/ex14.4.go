package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age} // u := &User{} or var u = new(User) 이런식으로 구조체 생성함 원래
	return &u
}
func main() {
	userPointer := NewUser("AAA", 123)

	fmt.Println(userPointer)
}
