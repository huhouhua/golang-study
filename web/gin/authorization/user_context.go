package authorization

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-study/web/gin/services"
)

type IUserContext interface {
	GetUserInfo() (*UserInfo, error)
	GetToken() string
}
type UserContext struct {
	ctx gin.Context
}

func NewUserProvider(ctx gin.Context) IUserContext {
	return &UserContext{
		ctx: ctx,
	}
}

func (u *UserContext) GetUserInfo() (*UserInfo, error) {
	userName := u.ctx.GetHeader("user")
	if userName == "" {
		return nil, errors.New("header 用户名字段为空！")
	}
	user, err := services.GetInfo(userName)
	if err != nil {
		return nil, err
	}
	return NewUserInfo(user.ID, user.Name, user.Password, u.GetToken()), nil
}

func (u *UserContext) GetToken() string {
	return u.ctx.GetHeader("token")
}
