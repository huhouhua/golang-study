//go:build e2e

package accesslog

import (
	"fmt"
	v3 "golang-study/web/v3"
	"testing"
)

func TestMiddlewareBuilderE2E(t *testing.T) {
	builder := MiddlewareBuilder{}
	mdl := builder.LogFunc(func(log string) {
		fmt.Println(log)

	}).Build()
	server := v3.NewHTTPServer(v3.ServerWithMiddleware(mdl))
	server.Get("/a/b/*", func(ctx *v3.Context) {
		ctx.Response.Write([]byte("hello,it's me "))
	})
	server.Start(":8081")

}
