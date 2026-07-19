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
	mux := http.NewServeMux() //먹스 사용, 라우터 역할
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) { // 절대경로/bar 경로
		fmt.Fprint(w, "Hello Bar!")
	})
	mux.Handle("/foo", &fooHandler{}) //인스턴스 형태로 등록.

	http.ListenAndServe(":3000", mux) //웹서버 구동, 리퀘스트 기다린다.
}
