package v3

import (
	"net"
	"net/http"
)

// 添加约束，确保HTTPServer类一定实现了Server接口
var _ Server = &HTTPServer{}

type HandlerFunc func(ctx *Context)

// Server 定义，根据其他情况，给Server做一次封装。
type Server interface {
	http.Handler
	//Start 服务启动
	Start(addr string) error

	//addRoute 添加路由注册，用来处理 请求处理
	addRoute(method string, path string, handlerFunc HandlerFunc)
}

type HTTPServer struct {
	*router
}

type HTTPSServer struct {
	HTTPServer
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}

func (h *HTTPServer) addRoute(method string, path string, handlerFunc HandlerFunc) {
	h.router.addRoute(method, path, handlerFunc)
}

func (h *HTTPServer) Get(method string, path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodGet, path, handlerFunc)
}
func (h *HTTPServer) Post(method string, path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodPost, path, handlerFunc)
}
func (h *HTTPServer) Delete(method string, path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodDelete, path, handlerFunc)
}
func (h *HTTPServer) Put(method string, path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodPut, path, handlerFunc)
}

// ServeHTTP 处理请求入口,是我们整个Web框架的核心入口，我们将在整个方法内部完成:
// Context 构建
// 路由匹配
// 执行业务逻辑
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Request:  request,
		Response: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	n, ok := h.findRouter(ctx.Request.Method, ctx.Request.URL.Path)
	if !ok || n.handler == nil {
		//没有找到此路由, 返回404
		ctx.Response.WriteHeader(http.StatusNotFound)
		_, _ = ctx.Response.Write([]byte("NOT FOUND"))
		return
	}
	n.handler(ctx)
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 可以在这里做一些前置动作，比如拦截器、判断等。
	return http.Serve(l, h)
}
