package user

import (
	"github.com/gin-gonic/gin"
	"golang-study/web/gin/services"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) UserInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, "Id 不能为空！")
		return
	}
	user, err := services.GetInfoById(id)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
