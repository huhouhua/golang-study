package web

import (
	"net"
	"net/http"
)

// 添加约束，确保HTTPServer类一定实现了Server接口
var _ Server = &HTTPServer{}

type HandlerFunc func(ctx Context)

// Server 定义，根据其他情况，给Server做一次封装。
type Server interface {
	http.Handler
	//Start 服务启动
	Start(addr string) error

	//AddRoute 添加路由注册，用来处理 请求处理
	AddRoute(method string, path string, handlerFunc HandlerFunc)
}

type HTTPServer struct {
}

type HTTPSServer struct {
	HTTPServer
}

func (H *HTTPServer) AddRoute(method string, path string, handlerFunc HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (H *HTTPServer) Get(method string, path string, handlerFunc HandlerFunc) {
	H.AddRoute(http.MethodGet, path, handlerFunc)
}
func (H *HTTPServer) Post(method string, path string, handlerFunc HandlerFunc) {
	H.AddRoute(http.MethodPost, path, handlerFunc)
}
func (H *HTTPServer) Delete(method string, path string, handlerFunc HandlerFunc) {
	H.AddRoute(http.MethodDelete, path, handlerFunc)
}
func (H *HTTPServer) Put(method string, path string, handlerFunc HandlerFunc) {
	H.AddRoute(http.MethodPut, path, handlerFunc)
}

// ServeHTTP 处理请求入口,是我们整个Web框架的核心入口，我们将在整个方法内部完成:
// Context 构建
// 路由匹配
// 执行业务逻辑
func (H *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Request:  request,
		Response: writer,
	}
	H.serve(ctx)
}

func (H *HTTPServer) serve(ctx *Context) {

}

func (H *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 可以在这里做一些前置动作，比如拦截器、判断等。
	return http.Serve(l, H)
}
