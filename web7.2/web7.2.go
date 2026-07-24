package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score,omitempty"`
}

var students map[int]Student //GET 예제라서 포인터 안쓴듯?
var lastId int

func SetupHandlers(g *gin.Engine) {
	g.GET("/students", GetStudentsHandler)
	g.GET("/students/:id", GetStudentHandler)
	g.POST("/student", PostStudentHandler)
	g.DELETE("/student/:id", DeleteStudentHandler)
	students = make(map[int]Student)
	students[1] = Student{Id: 1, Name: "aaa", Age: 16, Score: 87}
	students[2] = Student{Id: 2, Name: "bbb", Age: 18, Score: 98}
	lastId = 2
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

func GetStudentsHandler(c *gin.Context) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	c.JSON(http.StatusOK, list)
}
func GetStudentHandler(c *gin.Context) { //GET
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound) //에러코드만 보내기
		return
	}
	c.JSON(http.StatusOK, student) //json형태로 보내기
}
func PostStudentHandler(c *gin.Context) { // POST
	var student Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lastId++ //id는 서버에서 설정
	student.Id = lastId
	students[lastId] = student
	c.String(http.StatusCreated, "Success to add id:%d", lastId) //문자형태로 보내기
}
func DeleteStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	delete(students, id)
	c.String(http.StatusOK, "success to delete")
}

func main() {
	r := gin.Default()
	SetupHandlers(r)
	r.Run(":3000")
}
