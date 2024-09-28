package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func handler1(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "name1 is required")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "name2 is required")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
func main() {
	//配置路由
	http.HandleFunc("/1", handler1)
	http.HandleFunc("/2", handler2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
