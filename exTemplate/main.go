package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := &User{Name: "pjh", Email: "pjhsdf@naver.com", Age: 23}
	user2 := &User{Name: "aaa", Email: "aaa@daum.net", Age: 1}
	//tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}") // 틀 설정
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl") // 틀 설정, 방법 2. 틀을 파일로 빼놓고 구현

	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
}
