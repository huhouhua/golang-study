//go:build e2e

package opentelemetry

import (
	"go.opentelemetry.io/otel"
	v3 "golang-study/web/v3"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	tracer := otel.GetTracerProvider().Tracer(instrumentationName)
	builder := MiddlewareBuilder{
		Tracer: tracer,
	}
	server := v3.NewHTTPServer(v3.ServerWithMiddleware(builder.Build()))

	server.Get("/user", func(ctx *v3.Context) {
		c, span := tracer.Start(ctx.Request.Context(), "first_layer")
		defer span.End()
		c, second := tracer.Start(c, "second_layer")
		time.Sleep(time.Second)

		c, third := tracer.Start(c, "third_layer_1")
		time.Sleep(100 * time.Millisecond)
		third.End()

		c, third2 := tracer.Start(c, "third_layer_1")
		time.Sleep(300 * time.Millisecond)
		third2.End()

		second.End()

	})
	server.Start(":8081")

}
