package gin

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetUser(ctx *gin.Context) {
	ctx.String(200, "hello, world!")
}
func (c *UserController) GetJson(ctx *gin.Context) {
	user := &User{
		Id:   22,
		Name: "张三",
	}
	ctx.JSON(200, user)
}

type User struct {
	Id   int
	Name string
}
