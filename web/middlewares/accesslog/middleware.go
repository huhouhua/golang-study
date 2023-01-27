package accesslog

import (
	"encoding/json"
	v3 "golang-study/web/v3"
)

type MiddlewareBuilder struct {
	logFunc func(log string)
}

func (m *MiddlewareBuilder) LogFunc(fu func(log string)) *MiddlewareBuilder {
	m.logFunc = fu
	return m
}

func (m MiddlewareBuilder) Build() v3.Middleware {
	return func(next v3.HandlerFunc) v3.HandlerFunc {
		return func(ctx *v3.Context) {
			defer func() {
				l := accessLog{
					Host:       ctx.Request.Host,
					Route:      ctx.MatchedRoute,
					HTTPMethod: ctx.Request.Method,
					Path:       ctx.Request.URL.Path,
				}
				data, _ := json.Marshal(l)
				m.logFunc(string(data))
			}()
			next(ctx)
		}
	}
}

type accessLog struct {
	Host       string `json:"host,omitempty"`
	Route      string `json:"route,omitempty"`
	HTTPMethod string `json:"http_method,omitempty"`
	Path       string `json:"path,omitempty"`
}
