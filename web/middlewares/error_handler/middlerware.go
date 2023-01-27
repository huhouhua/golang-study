package error_handler

import v3 "golang-study/web/v3"

type MiddlewareBuilder struct {
	resp map[int][]byte
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		resp: map[int][]byte{},
	}
}
func (m *MiddlewareBuilder) AddCode(status int, data []byte) *MiddlewareBuilder {
	m.resp[status] = data
	return m
}

func (m *MiddlewareBuilder) Build() v3.Middleware {
	return func(next v3.HandlerFunc) v3.HandlerFunc {
		return func(ctx *v3.Context) {
			next(ctx)
			resp, ok := m.resp[ctx.ResponseStatusCode]
			if ok {
				//修改响应结果
				ctx.ResponseData = resp
			}
		}
	}
}
