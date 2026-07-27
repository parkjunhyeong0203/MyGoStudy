package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{Name: "pjh", Email: "pjh@daum.net"}

	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	/*w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))*/
	rd.JSON(w, http.StatusOK, user) //이거 한줄로 끝, 근데 gin에도 render있어서 그냥 gin쓰면 됨..
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	rd.HTML(w, http.StatusOK, "tmpl1", "pjh")
}
func main() {
	rd = render.New(render.Options{ // 확장자 추가해준다
		Extensions: []string{".html", "tmpl"},
	})
	mux := pat.New() //gorilla/mux 보다 쉬운 버전인 pat
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	n := negroni.Classic()
	n.UseHandler(mux) //decorater 처럼 래핑해서 부가기능 추가해준다. neg는 그런 패키지, log기능, 파일서버 기능등
	http.ListenAndServe(":3000", mux)
}
