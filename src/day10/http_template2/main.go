package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myTemplate *template.Template

type Person struct {
	Name string
	Title string
	age string
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	p := Person{Name: "chenyi", age: "18", Title: "我的个人网站"}
	myTemplate.Execute(w, p)
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err: ", err)
		return
	}
	return
}

func main() {
	initTemplate("C:/Workspace/go-study/src/day10/http_template2/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}