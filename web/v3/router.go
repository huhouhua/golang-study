package v3

import (
	"fmt"
	"strings"
)

// 用来支持对路由树的操作，做请求路径匹配
type router struct {
	// http method 到 路由树根节点
	trees map[string]*node
}

type node struct {
	router string

	path string

	//子path 到子节点的映射
	children map[string]*node

	//通匹符匹配
	starChild *node

	//路径参数
	paramChild *node

	//用户注册处理逻辑
	handler HandlerFunc
}

type matchInfo struct {
	n          *node
	pathParams map[string]string
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

// 创建子节点
func (n *node) childOfCreate(seg string) *node {
	//参数匹配
	if seg[0] == ':' {
		if n.starChild != nil {
			panic("web:不允许同时注册路径参数通配符匹配,已有通配符匹配！")
		}
		n.paramChild = &node{
			path: seg,
		}
		return n.paramChild
	}
	//通配符匹配
	if seg == "*" {
		if n.paramChild != nil {
			panic("web:不允许同时注册路径参数通配符匹配，已有路径参数匹配！")
		}

		n.starChild = &node{
			path: seg,
		}
		return n.starChild
	}
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[seg]
	if !ok {
		res = &node{
			path: seg,
		}
		n.children[seg] = res
	}
	return res
}

// childOf 优先匹配静态匹配，匹配不上，再考虑通配符匹配
// 第一个返回值是 子节点
// 第二个返回值是 标记是否是路径参数
// 第三个返回值是 有没有匹配到节点
func (n *node) childOf(path string) (*node, bool, bool) {
	if n.children == nil {

		if n.paramChild != nil {
			return n.paramChild, true, true
		}

		return n.starChild, false, n.starChild != nil
	}
	child, ok := n.children[path]
	if !ok {

		if n.paramChild != nil {
			return n.paramChild, true, true
		}
		return n.starChild, false, n.starChild != nil
	}
	return child, false, ok
}

// 查找路由匹配
func (r *router) findRouter(method string, path string) (*matchInfo, bool) {
	//找出对应的请求方式
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return &matchInfo{
			n: root,
		}, true
	}

	//把前缀和后缀的 / 都去掉
	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")

	var pathParams map[string]string

	for _, seg := range segs {
		child, isParam, found := root.childOf(seg)
		if !found {
			return nil, false
		}
		//命中了路径参数
		if isParam {
			if pathParams == nil {
				pathParams = make(map[string]string)
			}
			//path 是:id 这种形式
			pathParams[child.path[1:]] = seg
		}
		//找到了，赋值
		root = child
	}
	//找到了，返回节点
	return &matchInfo{
		n:          root,
		pathParams: pathParams,
	}, true
}

func (r *router) addRoute(method string, path string, handlerFunc HandlerFunc) {
	if path == "" {
		panic("path不能为空！")
	}
	//找到树
	root, ok := r.trees[method]
	if !ok {
		//没有根节点
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	if path[0] != '/' {
		panic("web:路径必须以 / 开头！")
	}
	if path != "/" && path[len(path)-1] == '/' {
		panic("web:路径必须以 / 结尾！")
	}

	//如果是根节点特殊处理下
	if path == "/" {
		//根节点重复注册
		if root.handler != nil {
			panic("web:路由冲突，重复注册[/]")
		}
		root.handler = handlerFunc
		root.router = "/"
		return
	}

	//切割path
	segs := strings.Split(path[1:], "/")
	for _, seg := range segs {
		if seg == "" {
			panic("web:不能有连续的 / ")
		}

		//递归下去，找到位置
		//如果有节点不存在，需要创建此节点
		child := root.childOfCreate(seg)
		root = child

	}
	if root.handler != nil {
		panic(fmt.Sprintf("web:路由冲突，重复注册[%s]", path))
	}
	root.handler = handlerFunc
	root.router = path
}
