package main

import (
	"clean-web/common-handler/good/httpx"
	"fmt"
	"net/http"
)

// 处理器函数
func userHandler(r *http.Request) (data any, err error) {
	return fmt.Sprintf("Hello user %s!", r.URL.Query().Get("name")), nil
}
func orderHandler(r *http.Request) (data any, err error) {
	return fmt.Sprintf("Hello order%s!", r.URL.Query().Get("name")), nil
}

func main() {
	// 配置路由
	httpx.HandleFunc("/users", userHandler)
	httpx.HandleFunc("/orders", orderHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
