package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestUserController(t *testing.T) {
	g := gin.Default()
	ctrl := NewUserController()

	g.GET("/json", ctrl.GetJson)
	g.GET("/user", ctrl.GetUser)
	//g.GET("/user", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "hello %s", "world")
	//})
	v1 := g.Group("/file")
	{
		v1.GET("/assets", func(context *gin.Context) {
			context.String(200, "this is file api !")
		})
	}

	_ = g.Run(":8082")
}
