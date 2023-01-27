package v3

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHTTPServer_ServeHTTP(t *testing.T) {
	server := NewHTTPServer()
	server.mdls = []Middleware{
		func(next HandlerFunc) HandlerFunc {
			return func(ctx *Context) {
				fmt.Println("第一个before")
				next(ctx)
				fmt.Println("第一个after")
			}
		},
		func(next HandlerFunc) HandlerFunc {
			return func(ctx *Context) {
				fmt.Println("第二个before")
				next(ctx)
				fmt.Println("第二个after")
			}
		},
		func(next HandlerFunc) HandlerFunc {
			return func(ctx *Context) {
				fmt.Println("第三个终断！")
			}
		},
		func(next HandlerFunc) HandlerFunc {
			return func(ctx *Context) {
				fmt.Println("第四个！")
			}
		},
	}
	server.ServeHTTP(nil, &http.Request{})
}
