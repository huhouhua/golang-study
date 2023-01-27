package v3

// Middleware 函数式的责任链设计
// 函数式的洋葱描述
type Middleware func(next HandlerFunc) HandlerFunc

// 非函数式设计
//type MiddlewareV1 interface {
//	Invoke(next HandlerFunc) HandlerFunc
//}

// 拦截器设计
//type Interceptor interface {
//	Before(ctx *Context)
//	After(ctx *Context)
//	Surround(ctx *Context)
//}

//集中式设计
//type Chain []HandlerFunc
//
//type HandlerFuncV1 func(ctx *Context) (next bool)
//
//type ChainV1 struct {
//	handlers []HandlerFuncV1
//}
//
//func (c ChainV1) Run(ctx *Context) {
//	for _, h := range c.handlers {
//		next := h(ctx)
//		if !next {
//			return
//		}
//	}
//}
