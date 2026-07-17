package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //Handler 인터페이스 구현
	fmt.Fprint(w, "Hello Foo!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //핸들러 등록, 경로에 리퀘스트 들어오면 핸들, 절대 경로
		fmt.Fprint(w, "Hello World") // 펑션은 정해진 형태의 인자를 받는다.
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) { // 절대경로/bar 경로
		fmt.Fprint(w, "Hello Bar!")
	})
	http.Handle("/foo", &fooHandler{}) //인스턴스 형태로 등록.

	http.ListenAndServe(":3000", nil) //웹서버 구동, 리퀘스트 기다린다.
}
