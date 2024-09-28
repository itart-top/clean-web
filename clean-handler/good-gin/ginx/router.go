package ginx

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
)

// 响应不处理返回改值
var Discard = struct{}{}

// 自定义成功处理函数。一般在启动时初始化: ginx.SuccessHandlerFunc = xx（注：线程不安全，只能初始化一次）
var SuccessHandlerFunc func(ctx *gin.Context, data any)

// 自定义错误处理函数。一般在启动时初始化: ginx.ErrHandlerFunc = xx（注：线程不安全，只能初始化一次）
var ErrHandlerFunc func(ctx *gin.Context, err error)

// 自定义handle.
type HandlerFunc func(ctx *gin.Context) (data any, err error)

type Router interface {
	GET(s string, h HandlerFunc)
	POST(s string, h HandlerFunc)
	DELETE(s string, h HandlerFunc)
	PUT(s string, h HandlerFunc)
	Group(relativePath string, handlers ...gin.HandlerFunc) Router
	RawRouter() gin.IRouter
	Use(middleware ...gin.HandlerFunc)
}

func WrapRouter(gr gin.IRouter) *router { // nolint:revive
	return &router{
		gr: gr,
	}
}

// 统一返回格式
type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// 路由封装.
type router struct {
	gr gin.IRouter
}

func (r router) Group(relativePath string, handlers ...gin.HandlerFunc) Router { // nolint
	return router{gr: r.gr.Group(relativePath, handlers...)}
}

func (r router) GET(s string, h HandlerFunc) {
	debugRouterPrint(r.gr, http.MethodGet, s, h)
	r.gr.GET(s, r.handlerAdapt(h))
}

func (r router) POST(s string, h HandlerFunc) {
	debugRouterPrint(r.gr, http.MethodPost, s, h)
	r.gr.POST(s, r.handlerAdapt(h))
}

func (r router) DELETE(s string, h HandlerFunc) {
	debugRouterPrint(r.gr, http.MethodDelete, s, h)
	r.gr.DELETE(s, r.handlerAdapt(h))
}

func (r router) PUT(s string, h HandlerFunc) {
	debugRouterPrint(r.gr, http.MethodPut, s, h)
	r.gr.PUT(s, r.handlerAdapt(h))
}

func (r router) RawRouter() gin.IRouter {
	return r.gr
}

func (r router) Use(middleware ...gin.HandlerFunc) {
	r.gr.Use(middleware...)
}

func (r router) handlerAdapt(h HandlerFunc) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			r.errHandler(ctx, err)
		} else {
			r.successHandler(ctx, data)
		}
	}
}

func (r router) successHandler(ctx *gin.Context, data any) {
	if data == Discard {
		return
	}
	if SuccessHandlerFunc != nil {
		SuccessHandlerFunc(ctx, data)
		return
	}

	resp := Response{
		Code:    200,
		Success: true,
		Msg:     "success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (r router) errHandler(ctx *gin.Context, err error) {
	if ErrHandlerFunc != nil {
		ErrHandlerFunc(ctx, err)
	} else {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}

func nameOfFunction(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
