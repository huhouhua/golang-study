package error_handler

import (
	v3 "golang-study/web/v3"
	"net/http"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := NewMiddlewareBuilder()

	notFound := []byte(`
	 <html>
		<body>
			<h1 style="color:#eb1616"> 这个页面不存在！</h1>
		</body>
	</html>`)
	badRequest := []byte(`
	<html>
		<body>
			<h1 style="color:#172bcf"> 请求不对！</h1>
		</body>
	</html>`)

	builder.AddCode(http.StatusNotFound, notFound).
		AddCode(http.StatusBadRequest, badRequest)

	server := v3.NewHTTPServer(v3.ServerWithMiddleware(builder.Build()))
	server.Start(":8081")

}
