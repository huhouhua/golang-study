package recover

import (
	"fmt"
	v3 "golang-study/web/v3"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := &MiddlewareBuilder{
		StatusCode: 500,
		Data:       []byte("你好 panic了"),
		Log: func(ctx *v3.Context) {
			fmt.Printf("panic路径:%s", ctx.Request.URL.String())
		},
	}
	server := v3.NewHTTPServer(v3.ServerWithMiddleware(builder.Build()))
	server.Get("/user", func(ctx *v3.Context) {
		panic("方发生panic了")
	})
	server.Start(":8081")
}
