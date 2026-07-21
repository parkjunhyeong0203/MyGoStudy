package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 기본 라우트 설정
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Secure WSL2 HTTPS Server!")
	})

	// ListenAndServeTLS를 사용하여 HTTPS 서버 실행
	err := http.ListenAndServeTLS(":3000", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
