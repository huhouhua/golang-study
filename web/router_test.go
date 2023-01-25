package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_AddRoute(t *testing.T) {
	testRoutes := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/user/home",
		},
	}

	var mockHandler HandlerFunc = func(ctx Context) {

	}

	r := newRouter()
	for _, route := range testRoutes {
		r.AddRoute(route.method, route.path, mockHandler)
	}

	wantRouter := &router{
		trees: map[string]*node{
			http.MethodGet: &node{
				path: "/",
				children: map[string]*node{
					"user": &node{
						path: "user",
						children: map[string]*node{
							"home": &node{
								path:    "home",
								handler: mockHandler,
							},
						},
					},
				},
			},
		},
	}
	msg, ok := wantRouter.equal(r)
	assert.True(t, ok, msg)
}

// string 返回的是一个错误信息，帮助我们排查问题
// bool 是代表是否真的相等。

func (r *router) equal(y *router) (string, bool) {
	for k, v := range r.trees {

		dst, ok := y.trees[k]
		if !ok {
			return fmt.Sprintf("找不到对应的Http method "), false
		}

		msg, equal := v.equal(dst)
		if !equal {
			return msg, false
		}

	}
	return "", true
}

func (n *node) equal(y *node) (string, bool) {
	if n.path != y.path {
		return fmt.Sprintf("节点路径不匹配"), false
	}
	if len(n.children) != len(y.children) {
		return fmt.Sprintf("子节点数量不相等！"), false
	}

	nhandler := reflect.ValueOf(n.handler)
	yhandler := reflect.ValueOf(y.handler)
	if nhandler != yhandler {
		return fmt.Sprintf("handler 不相等！"), false
	}

	for path, c := range n.children {

		dst, ok := y.children[path]
		if !ok {
			return fmt.Sprintf("子节点 %s 不存在", path), false
		}

		msg, ok := c.equal(dst)
		if !ok {
			return msg, false
		}

	}
	return "", true
}