package opentelemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	v3 "golang-study/web/v3"
)

const instrumentationName = "golang-study/web/v3/middlewares/opentelemetry"

type MiddlewareBuilder struct {
	Tracer trace.Tracer
}

func (m MiddlewareBuilder) Build() v3.Middleware {
	if m.Tracer == nil {
		m.Tracer = otel.GetTracerProvider().Tracer(instrumentationName)
	}
	return func(next v3.HandlerFunc) v3.HandlerFunc {
		return func(ctx *v3.Context) {

			reqCtx := ctx.Request.Context()

			//尝试和客户端的trace结合在一起
			reqCtx = otel.GetTextMapPropagator().Extract(reqCtx, propagation.HeaderCarrier(ctx.Request.Header))

			reqCtx, span := m.Tracer.Start(reqCtx, "unknown")
			defer span.End()

			//按照自己的情况，记录需要的请求数据
			span.SetAttributes(attribute.String("http.method", ctx.Request.Method))
			span.SetAttributes(attribute.String("http.url", ctx.Request.URL.String()))
			span.SetAttributes(attribute.String("http.scheme", ctx.Request.URL.Scheme))
			span.SetAttributes(attribute.String("http.host", ctx.Request.Host))

			ctx.Request = ctx.Request.WithContext(reqCtx)
			//ctx.Ctx = reqCtx

			next(ctx)

			// 最后执行完，设置路由，当做链路名称
			span.SetName(ctx.MatchedRoute)

			span.SetAttributes(attribute.Int("http.status", ctx.ResponseStatusCode))
		}
	}
}
