package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

var mockHandler HandlerFunc = func(ctx Context) {

}

func TestRouter_AddRoute(t *testing.T) {
	caseRouters := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/",
		},
		{
			method: http.MethodGet,
			path:   "/user",
		},
		{
			method: http.MethodGet,
			path:   "/user/home",
		},
		{
			method: http.MethodGet,
			path:   "/order/detail",
		},
		{
			method: http.MethodPost,
			path:   "/",
		},
		{
			method: http.MethodPost,
			path:   "/order/create",
		},
		{
			method: http.MethodPost,
			path:   "/login",
		},
	}
	r := newRouter()
	for _, route := range caseRouters {
		r.AddRoute(route.method, route.path, mockHandler)
	}

	wantRouter := &router{
		trees: map[string]*node{
			http.MethodGet: &node{
				path:    "/",
				handler: mockHandler,
				children: map[string]*node{
					"user": &node{
						path:    "user",
						handler: mockHandler,
						children: map[string]*node{
							"home": &node{
								path:    "home",
								handler: mockHandler,
							},
						},
					},
					"order": &node{
						path: "order",
						children: map[string]*node{
							"detail": &node{
								path:    "detail",
								handler: mockHandler,
							},
						},
					},
				},
			},
			http.MethodPost: &node{
				path:    "/",
				handler: mockHandler,
				children: map[string]*node{
					"order": &node{
						path: "order",
						children: map[string]*node{
							"create": &node{
								path:    "create",
								handler: mockHandler,
							},
						},
					},
					"login": &node{
						path:    "login",
						handler: mockHandler,
					},
				},
			},
		},
	}
	msg, ok := wantRouter.equal(r)
	assert.True(t, ok, msg)
}

// 路径校验测试
func TestRouter_Path_Validation(t *testing.T) {
	r := newRouter()

	assert.Panicsf(t, func() {
		r.AddRoute(http.MethodGet, "", mockHandler)
	}, "path不能为空！")

	assert.Panicsf(t, func() {
		r.AddRoute(http.MethodGet, "a/", mockHandler)
	}, "web:路径必须以 / 开头！")

	assert.Panicsf(t, func() {
		r.AddRoute(http.MethodGet, "/a/b/c/", mockHandler)
	}, "web:路径必须以 / 结尾！")

	assert.Panicsf(t, func() {
		r.AddRoute(http.MethodGet, "/a//b/", mockHandler)
	}, "web:不能有连续的 / ")

}

// 路径重复校验测试
func TestRouter_Path_Repetition(t *testing.T) {
	r := newRouter()
	r.AddRoute(http.MethodGet, "/", mockHandler)
	assert.Panicsf(t, func() {
		r.AddRoute(http.MethodGet, "/", mockHandler)
	}, "web:路由冲突，重复注册[/]")

	r.AddRoute(http.MethodGet, "/a/b/c", mockHandler)
	assert.Panicsf(t, func() {
		r.AddRoute(http.MethodGet, "/a/b/c", mockHandler)
	}, "web:路由冲突，重复注册[/a/b/c]")
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
