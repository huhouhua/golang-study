//go:build e2e

package v3

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	h := NewHTTPServer()
	h.Get("/user", func(ctx *Context) {

	})
	h.Get("/order/detail", func(ctx *Context) {
		ctx.Response.Write([]byte("hello,order detail!"))
	})
	h.Get("/order/abc", func(ctx *Context) {
		ctx.Response.Write([]byte(fmt.Sprintf("hello,%s", ctx.Request.URL.Path)))
	})
	h.Get("/order/*/ass", func(ctx *Context) {
		ctx.Response.Write([]byte(fmt.Sprintf("hello 通配符,%s", ctx.Request.URL.Path)))
	})
	h.Get("/order/*/*/ass", func(ctx *Context) {

		ctx.Response.Write([]byte(fmt.Sprintf("hello 通配符,%s", ctx.Request.URL.Path)))
	})
	h.Get("/user/hu", func(ctx *Context) {
		ctx.RespJSON(http.StatusOK, User{
			Name: "huhouhua",
		})
	})
	err := h.Start(":8081")
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name string `json:"name"`
}
