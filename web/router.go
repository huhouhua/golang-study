package web

// 用来支持对路由树的操作，做请求路径匹配
type router struct {
	// http method 到 路由树根节点
	trees map[string]*node
}

type node struct {
	path string

	//子path 到子节点的映射
	children map[string]*node

	//用户注册处理逻辑
	handler HandlerFunc
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

func (r *router) AddRoute(method string, path string, handlerFunc HandlerFunc) {

}
