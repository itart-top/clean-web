package ginx

import (
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// Fix "[GIN-debug]"默认都打印 "baishancloud/gaea/internal/pkg/ginx.router.handlerAdapt"问题
func debugRouterPrint(router gin.IRouter, httpMethod string, absolutePath string, handler HandlerFunc) {

	if !gin.IsDebugging() {
		return
	}
	if rg, ok := router.(interface {
		BasePath() string
	}); ok {
		absolutePath = path.Join(strings.TrimRight(rg.BasePath(), "/"), strings.TrimLeft(absolutePath, "/"))
	}
	handlerName := nameOfFunction(handler)
	format := "[GINX-debug] %-6s %-25s --> %s\n"
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(gin.DefaultWriter, format, httpMethod, absolutePath, handlerName)
}
