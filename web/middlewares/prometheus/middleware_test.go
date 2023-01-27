//go:build e2e

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v3 "golang-study/web/v3"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := MiddlewareBuilder{
		Namespace: "alibaba",
		Subsystem: "web",
		Name:      "http_response",
	}

	server := v3.NewHTTPServer(v3.ServerWithMiddleware(builder.Build()))
	server.Get("/user", func(ctx *v3.Context) {

		val := rand.Intn(1000) + 1
		time.Sleep(time.Duration(val) * time.Millisecond)

		ctx.RespJSON(202, User{
			Name: "huhouhua",
		})
	})
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8082", nil)
	}()
	server.Start(":8081")
}

type User struct {
	Name string
}
