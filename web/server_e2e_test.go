// go:build

package web

import (
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	h := &HTTPServer{}
	h.AddRoute(http.MethodGet, "/user", func(ctx Context) {

	})

	err := h.Start(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
