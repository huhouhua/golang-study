package v3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

//func GenerateTest_With_Asterisk() {
//
//}

var mockHandler HandlerFunc = func(ctx *Context) {

}

func TestRouter_AddRoute(t *testing.T) {
	testRouters := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/",
		},
		{
			method: http.MethodGet,
			path:   "/*",
		},
		{
			method: http.MethodGet,
			path:   "/*/*",
		},
		{
			method: http.MethodGet,
			path:   "/*/abc",
		},
		{
			method: http.MethodGet,
			path:   "/*/abc/ddd",
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
			method: http.MethodGet,
			path:   "/order/*",
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
	for _, route := range testRouters {
		r.addRoute(route.method, route.path, mockHandler)
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
						starChild: &node{
							path:    "*",
							handler: mockHandler,
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
		r.addRoute(http.MethodGet, "", mockHandler)
	}, "path不能为空！")

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "a/", mockHandler)
	}, "web:路径必须以 / 开头！")

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a/b/c/", mockHandler)
	}, "web:路径必须以 / 结尾！")

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a//b/", mockHandler)
	}, "web:不能有连续的 / ")

}

// 路径重复校验测试
func TestRouter_Path_Repetition(t *testing.T) {
	r := newRouter()
	r.addRoute(http.MethodGet, "/", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/", mockHandler)
	}, "web:路由冲突，重复注册[/]")

	r.addRoute(http.MethodGet, "/a/b/c", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a/b/c", mockHandler)
	}, "web:路由冲突，重复注册[/a/b/c]")
}

// 测试查找路由
func TestRouter_findRoute(t *testing.T) {
	testRoutes := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodDelete,
			path:   "/",
		},
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
			method: http.MethodGet,
			path:   "/order/*",
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
	for _, route := range testRoutes {
		r.addRoute(route.method, route.path, mockHandler)
	}
	testCases := []struct {
		name        string
		method      string
		path        string
		description string
		wantFound   bool
		wantNode    *node
	}{
		{
			name:        "method no found",
			method:      http.MethodOptions,
			path:        "/order/detail",
			description: "方法都不存在",
			wantFound:   false,
		},
		{
			name:        "path no found",
			method:      http.MethodGet,
			path:        "/wadasdsad",
			description: "路径不存在！",
			wantFound:   false,
		},
		{
			name:        "order detail",
			method:      http.MethodGet,
			path:        "/order/detail",
			description: "完全命中",
			wantFound:   true,
			wantNode: &node{
				handler: mockHandler,
				path:    "detail",
			},
		},
		{
			name:        "order start",
			method:      http.MethodGet,
			path:        "/order/abc",
			description: "通配符匹配",
			wantFound:   true,
			wantNode: &node{
				handler: mockHandler,
				path:    "*",
			},
		},
		{
			name:        "order",
			method:      http.MethodGet,
			path:        "/order",
			description: "命中了，但是没有Handler",
			wantFound:   true,
			wantNode: &node{
				//handler: mockHandler,
				path: "order",
				children: map[string]*node{
					"detail": &node{
						handler: mockHandler,
						path:    "detail",
					},
				},
			},
		},
		{
			name:        "root",
			method:      http.MethodDelete,
			path:        "/",
			description: "根节点",
			wantFound:   true,
			wantNode: &node{
				path:    "/",
				handler: mockHandler,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("描述:%s \n", tc.description)

			n, found := r.findRouter(tc.method, tc.path)
			assert.Equal(t, tc.wantFound, found)
			if !found {
				return
			}
			assert.Equal(t, tc.wantNode.path, n.path)

			msg, ok := tc.wantNode.equal(n)
			assert.True(t, ok, msg)

			msg, ok = equalHandler(tc.wantNode.handler, n.handler)
			assert.True(t, ok, msg)
		})
	}

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
	if n.starChild != nil {
		msg, ok := n.starChild.equal(y.starChild)
		if !ok {
			return msg, ok
		}
	}

	if msg, ok := equalHandler(n.handler, y.handler); !ok {
		return msg, ok
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

// 比较处理器是否相等
func equalHandler(n HandlerFunc, y HandlerFunc) (string, bool) {
	nHandler := reflect.ValueOf(n)
	yHandler := reflect.ValueOf(y)
	if nHandler != yHandler {
		return fmt.Sprintf("handler 不相等！"), false
	}
	return "", true
}
