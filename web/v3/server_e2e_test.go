//go:build e2e

package v3

import (
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	h := NewHTTPServer()
	h.addRoute(http.MethodGet, "/user", func(ctx *Context) {

	})
	h.addRoute(http.MethodGet, "/order/detail", func(ctx *Context) {
		ctx.Response.Write([]byte("hello,order detail!"))
	})

	err := h.Start(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
