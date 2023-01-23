package iris

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"testing"
)

func TestIris(t *testing.T) {
	app := iris.New()
	app.Get("/", func(c context.Context) {
		_, _ = c.HTML("Hello <strong>%s</strong>!", "World")
	})
	_ = app.Listen(":8083")
}
