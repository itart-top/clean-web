package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler不再需要调用ResponseWriter方法，交由中间层处理
// data：正常响应结果
// err：异常时返回的错误信息
type HandlerFunc func(req *http.Request) (data any, err error)

type response struct {
	Code    int
	Message string
	Data    interface{}
}

// 统一处理异常，适配http.HandlerFunc
func HandlerAdapt(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := fn(req)
		if err == nil {
			successHandler(w, data)
		} else {
			errHandler(w, err)
		}

	}
}

func successHandler(w http.ResponseWriter, data any) {
	resp := response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	w.WriteHeader(http.StatusOK)
	// 响应结果处理
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Printf("json encode error: %v", err)
	}

}

func errHandler(w http.ResponseWriter, err error) {
	resp := response{
		Code:    500,
		Message: "error",
		Data:    err.Error(),
	}
	w.WriteHeader(http.StatusInternalServerError)
	// 响应结果处理
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Printf("json encode error: %v", err)
	}
}
