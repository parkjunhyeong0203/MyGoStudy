package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 자동으로 실행, 사라지지 않도록 _
	"time"
)

func main() {
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()
	select {
	case status := <-ch:
		w.WriteHeader(status)

	case <-time.After(200 * time.Millisecond):
		w.WriteHeader(http.StatusRequestTimeout)
	}

}
