package web

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

// 创建子节点
func (n *node) childOfCreate(seg string) *node {
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

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	child, ok := n.children[path]
	return child, ok
}

func (r *router) findRouter(method string, path string) (*node, bool) {
	//找出对应的请求方式
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}

	if path == "/" {
		return root, true
	}

	//把前缀和后缀的 / 都去掉
	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		child, found := root.childOf(seg)
		if !found {
			return nil, false
		}
		//找到了，赋值
		root = child
	}
	//找到了，返回节点
	return root, true
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
}
