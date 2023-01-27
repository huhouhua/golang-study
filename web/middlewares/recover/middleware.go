package recover

import v3 "golang-study/web/v3"

type MiddlewareBuilder struct {
	StatusCode int
	Data       []byte
	Log        func(ctx *v3.Context)
}

func (m *MiddlewareBuilder) Build() v3.Middleware {
	return func(next v3.HandlerFunc) v3.HandlerFunc {
		return func(ctx *v3.Context) {
			defer func() {
				if err := recover(); err != nil {
					ctx.ResponseData = m.Data
					ctx.ResponseStatusCode = m.StatusCode
					m.Log(ctx)
				}
			}()
			next(ctx)
		}
	}
}
