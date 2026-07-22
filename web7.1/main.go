package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

type Students []Student

func (s Students) Len() int { //Students는 sort인터페이스 구현
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

var students map[int]Student //GET 예제라서 포인터 안쓴듯?
var lastId int

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) { //GET, r은 안쓰고 내부 코드에서 전역변수로 객체 만들어서 사용
	list := make(Students, 0) // 어차피 GET 예제니까 r은 안쓴다.
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
func GetStudentHandler(w http.ResponseWriter, r *http.Request) { //GET
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func PostStudentHandler(w http.ResponseWriter, r *http.Request) { // POST, r로 들어온거 등록
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++ //id는 서버에서 설정
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated) //추가됐음을 알린다.
}
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK)
}

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
