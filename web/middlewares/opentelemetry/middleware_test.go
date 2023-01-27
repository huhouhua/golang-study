//go:build e2e

package opentelemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	v3 "golang-study/web/v3"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	tracer := otel.GetTracerProvider().Tracer(instrumentationName)
	builder := MiddlewareBuilder{
		Tracer: tracer,
	}
	server := v3.NewHTTPServer(v3.ServerWithMiddleware(builder.Build()))

	server.Get("/user/:name", func(ctx *v3.Context) {
		//最外面一层
		c, span := tracer.Start(ctx.Request.Context(), "first_layer")
		defer span.End()

		//第二层
		secondC, second := tracer.Start(c, "second_layer")
		time.Sleep(time.Second)

		//第三层
		_, third1 := tracer.Start(secondC, "third_layer_1")
		time.Sleep(100 * time.Millisecond)
		third1.End()

		//第四层
		_, third2 := tracer.Start(secondC, "third_layer_2")
		time.Sleep(300 * time.Millisecond)
		third2.End()

		//最外面一层介绍
		second.End()

		_, first := tracer.Start(ctx.Request.Context(), "third_layer_1")
		defer first.End()
		time.Sleep(100 * time.Millisecond)

		str := ctx.PathValue("name")

		ctx.RespJSON(
			http.StatusAccepted,
			User{
				Name: str.Val,
			})
		ctx.Response.Write([]byte("hello,world"))
	})
	initZipkin(t)

	server.Start(":8081")

}

// initZipkin 打开 http://localhost:19411 查看调用链路情况
func initZipkin(t *testing.T) {
	exporter, err := zipkin.New(
		"http://localhost:19411/api/v2/spans",
		zipkin.WithLogger(log.New(os.Stderr, "opentelemetry-demo", log.Ldate|log.Ltime|log.Llongfile)),
	)
	if err != nil {
		t.Fatal(err)
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("opentelemetry-demo"),
		)),
	)
	otel.SetTracerProvider(tp)
}

// initJeager 打开 http://localhost:14268 查看调用链路情况
func initJeager(t *testing.T) {
	url := "http://localhost:14268/api/traces"
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		t.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		// Record information about this application in a Resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("opentelemetry-demo"),
			attribute.String("environment", "dev"),
			attribute.Int64("ID", 1),
		)),
	)

	otel.SetTracerProvider(tp)
}

type User struct {
	Name string
}
