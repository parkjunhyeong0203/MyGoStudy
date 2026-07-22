package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"` //go의 어노테이션. go에서의 변수가 json에서 어떻게 쓰이는지 설명을 붙여준다.
	LastName  string    `json:"last_name"`
	Email     string    `json:"Email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //Handler 인터페이스 구현
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user) //리퀘스트 바디에 NewDecoder의 인자로 Reader를 받는다, 받은 json 디코딩 -> 그 후 user에 넣는다.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)                      // user를 다시 json
	w.Header().Add("content-type", "application/json") //이걸 알려줘야 json을 해석가능
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) //data는 바이트 타입이라 강제 타입 변환
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}
func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() //먹스 사용, 라우터 역할
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) { // 절대경로/bar 경로
		fmt.Fprint(w, "Hello Bar!")
	})
	mux.Handle("/foo", &fooHandler{}) //인스턴스 형태로 등록.
	return mux
}
