package httpx

import "net/http"

func HandleFunc(pattern string, handler HandlerFunc) {
	http.HandleFunc("/users", HandlerAdapt(handler))
}
