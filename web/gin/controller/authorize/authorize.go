package authorize

import (
	"github.com/gin-gonic/gin"
	"golang-study/web/gin/authorization"
	"golang-study/web/gin/services"
	"net/http"
)

type AuthzController struct {
}

func NewAuthzController() *AuthzController {
	return &AuthzController{}
}

func (a *AuthzController) Login(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, "用户名和密码不能为空！")
		return
	}
	user, err := services.GetInfo(username)
	if err != nil {
		c.JSON(http.StatusOK, "该用户不存在！")
		return
	}
	if !authorization.CheckPassword(password, user.ID, user.Password) {
		c.JSON(http.StatusOK, "密码不正确！")
		return
	}
	token, err := authorization.GeneralJwtToken(username)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func (a AuthzController) LoginOut(c *gin.Context) {
	//todo list
	//从Redis 删除key  即可

}
