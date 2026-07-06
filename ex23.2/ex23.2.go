package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "Password is short."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// return fmt.Errorf("...")
		// return error.New("...")
		return PasswordError{len(password), 8} //리턴했을때 객체 내부에 Error()함수가 있으면 %v 했을때 그거 실행
	}
	return nil
}
func main() {
	err := RegisterAccount("myId", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d ReauireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("success")
	}
}
