package v3

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
			path:   "/order/detail/:id",
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
								paramChild: &node{
									path:    ":id",
									handler: mockHandler,
								},
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

// ??????????????????
func TestRouter_Path_Validation(t *testing.T) {
	r := newRouter()

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "", mockHandler)
	}, "path???????????????")

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "a/", mockHandler)
	}, "web:??????????????? / ?????????")

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a/b/c/", mockHandler)
	}, "web:??????????????? / ?????????")

	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a//b/", mockHandler)
	}, "web:?????????????????? / ")

}

// ????????????????????????
func TestRouter_Path_Repetition(t *testing.T) {
	r := newRouter()
	r.addRoute(http.MethodGet, "/", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/", mockHandler)
	}, "web:???????????????????????????[/]")

	r.addRoute(http.MethodGet, "/a/b/c", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a/b/c", mockHandler)
	}, "web:???????????????????????????[/a/b/c]")
}

// ????????????????????????
func TestRoute_Path_Conflict(t *testing.T) {
	r := newRouter()
	r.addRoute(http.MethodGet, "/a/*", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a/:id", mockHandler)
	}, "web:????????????????????????????????????????????????,????????????????????????")

	r = newRouter()
	r.addRoute(http.MethodGet, "/a/:id", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a/*", mockHandler)
	}, "web:??????????????????????????????????????????????????????????????????????????????")
}

// ??????????????????
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
		{
			method: http.MethodPost,
			path:   "/login/:username",
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
		info        *matchInfo
	}{
		{
			name:        "method no found",
			method:      http.MethodOptions,
			path:        "/order/detail",
			description: "??????????????????",
			wantFound:   false,
		},
		{
			name:        "path no found",
			method:      http.MethodGet,
			path:        "/wadasdsad",
			description: "??????????????????",
			wantFound:   false,
		},
		{
			name:        "order detail",
			method:      http.MethodGet,
			path:        "/order/detail",
			description: "????????????",
			wantFound:   true,
			info: &matchInfo{
				n: &node{
					handler: mockHandler,
					path:    "detail",
				},
			},
		},
		{
			name:        "order start",
			method:      http.MethodGet,
			path:        "/order/abc",
			description: "???????????????",
			wantFound:   true,
			info: &matchInfo{
				n: &node{
					handler: mockHandler,
					path:    "*",
				},
			},
		},
		{
			name:        "order",
			method:      http.MethodGet,
			path:        "/order",
			description: "????????????????????????Handler",
			wantFound:   true,
			info: &matchInfo{
				n: &node{
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
		},
		{
			name:        "root",
			method:      http.MethodDelete,
			path:        "/",
			description: "?????????",
			wantFound:   true,
			info: &matchInfo{
				n: &node{
					path:    "/",
					handler: mockHandler,
				},
			},
		},
		{
			name:        "login username",
			method:      http.MethodPost,
			path:        "/login/huhouhua",
			description: "??????????????????",
			wantFound:   true,
			info: &matchInfo{
				n: &node{
					path:    ":username",
					handler: mockHandler,
				},
				pathParams: map[string]string{
					"username": "huhouhua",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("??????:%s \n", tc.description)

			n, found := r.findRouter(tc.method, tc.path)
			assert.Equal(t, tc.wantFound, found)
			if !found {
				return
			}
			assert.Equal(t, tc.info.n.path, n.n.path)

			assert.Equal(t, tc.info.pathParams, n.pathParams)
			msg, ok := tc.info.n.equal(n.n)
			assert.True(t, ok, msg)

			msg, ok = equalHandler(tc.info.n.handler, n.n.handler)
			assert.True(t, ok, msg)

		})
	}

}

// string ?????????????????????????????????????????????????????????
// bool ??????????????????????????????
func (r *router) equal(y *router) (string, bool) {
	for k, v := range r.trees {

		dst, ok := y.trees[k]
		if !ok {
			return fmt.Sprintf("??????????????????Http method "), false
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
		return fmt.Sprintf("?????????????????????"), false
	}
	if len(n.children) != len(y.children) {
		return fmt.Sprintf("???????????????????????????"), false
	}
	if n.starChild != nil {
		msg, ok := n.starChild.equal(y.starChild)
		if !ok {
			return msg, ok
		}
	}
	if n.paramChild != nil {
		msg, ok := n.paramChild.equal(y.paramChild)
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
			return fmt.Sprintf("????????? %s ?????????", path), false
		}

		msg, ok := c.equal(dst)
		if !ok {
			return msg, false
		}

	}
	return "", true
}

// ???????????????????????????
func equalHandler(n HandlerFunc, y HandlerFunc) (string, bool) {
	nHandler := reflect.ValueOf(n)
	yHandler := reflect.ValueOf(y)
	if nHandler != yHandler {
		return fmt.Sprintf("handler ????????????"), false
	}
	return "", true
}
