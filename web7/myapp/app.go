package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello user")
}
func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Hello user id:", vars["id"])
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler) //id마다 메서드를 하나하나 추가 할수 없으니 고릴라 먹스사용

	return mux
}
