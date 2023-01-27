package v3

import (
	"fmt"
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

	//中间件处理
	mdls []Middleware

	log func(msg string, args ...any)
}

type HTTPServerOption func(server *HTTPServer)

type HTTPSServer struct {
	HTTPServer
}

func NewHTTPServer(opts ...HTTPServerOption) *HTTPServer {
	res := &HTTPServer{
		router: newRouter(),
		log: func(msg string, args ...any) {
			fmt.Printf(msg, args...)
		},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func ServerWithMiddleware(mdls ...Middleware) HTTPServerOption {
	return func(server *HTTPServer) {
		server.mdls = mdls
	}
}

func (h *HTTPServer) addRoute(method string, path string, handlerFunc HandlerFunc) {
	h.router.addRoute(method, path, handlerFunc)
}

func (h *HTTPServer) Get(path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodGet, path, handlerFunc)
}
func (h *HTTPServer) Post(path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodPost, path, handlerFunc)
}
func (h *HTTPServer) Delete(path string, handlerFunc HandlerFunc) {
	h.addRoute(http.MethodDelete, path, handlerFunc)
}
func (h *HTTPServer) Put(path string, handlerFunc HandlerFunc) {
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
	// 最后一个是这个
	root := h.serve

	//然后这里就是利用最后一个不断往前面组装链条
	//从后往前
	//把后一个作为前一个的next 构造好链条
	for i := len(h.mdls) - 1; i >= 0; i-- {
		root = h.mdls[i](root)
	}

	var m Middleware = func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) {
			next(ctx)

			//最后设置到Response上  把ResponseStatusCode、ResponseData
			h.flashResponse(ctx)
		}
	}
	//这里执行的时候，就是从前往后
	root = m(root)
	root(ctx)
}

func (h *HTTPServer) flashResponse(ctx *Context) {
	if ctx.ResponseStatusCode != 0 {
		ctx.Response.WriteHeader(ctx.ResponseStatusCode)
	}
	n, err := ctx.Response.Write(ctx.ResponseData)
	if err != nil || n != len(ctx.ResponseData) {
		h.log("写入响应失败 %v", err)
	}
}

func (h *HTTPServer) serve(ctx *Context) {
	info, ok := h.findRouter(ctx.Request.Method, ctx.Request.URL.Path)
	if !ok || info.n.handler == nil {
		//没有找到此路由, 返回404
		ctx.ResponseStatusCode = http.StatusNotFound
		ctx.ResponseData = []byte("NOT FOUND")
		//ctx.Response.WriteHeader(http.StatusNotFound)
		//_, _ = ctx.Response.Write([]byte("NOT FOUND"))
		return
	}
	ctx.PathParams = info.pathParams
	ctx.MatchedRoute = info.n.router
	info.n.handler(ctx)
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 可以在这里做一些前置动作，比如拦截器、判断等。
	return http.Serve(l, h)
}
