package main

import (
	"clean-web/clean-handler/good-gin/ginx"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router := ginx.WrapRouter(engine)
	router.GET("/users", userHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Sprintf("ListenAndServe: %+v", err))
	}
}

// 处理hello请求的handler。如果有异常返回，响应结果也是直接放回
func userHandler(r *gin.Context) (data any, err error) {
	return fmt.Sprintf("Hello user %s!", r.Param("name")), nil
}
