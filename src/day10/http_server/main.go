package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello ")
}

func User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle User")
	fmt.Fprintf(w, "User ")
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user", User)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("http listen failed.")
	}
}
