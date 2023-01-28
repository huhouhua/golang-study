package accesslog

import (
	"fmt"
	v3 "golang-study/web/v3"
	"net/http"
	"testing"
)

func TestMiddlewareBuilder(t *testing.T) {
	builder := MiddlewareBuilder{}
	mdl := builder.LogFunc(func(log string) {
		fmt.Println(log)

	}).Build()
	server := v3.NewHTTPServer(v3.ServerWithMiddleware(mdl))
	server.Post("/a/b/*", func(ctx *v3.Context) {
		fmt.Println("hello,it's me")
	})
	req, err := http.NewRequest(http.MethodPost, "/a/b/c", nil)
	req.URL.Host = "localhost"
	if err != nil {
		t.Fatal(err)
	}
	server.ServeHTTP(nil, req)
}
