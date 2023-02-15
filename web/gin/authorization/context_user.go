package authorization

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-study/web/gin/services"
)

type IContextUser interface {
	GetUserInfo() (*UserInfo, error)
	GetToken() string
}
type ContextUser struct {
	ctx gin.Context
}

func NewUserProvider(ctx gin.Context) IContextUser {
	return &ContextUser{
		ctx: ctx,
	}
}

func (u *ContextUser) GetUserInfo() (*UserInfo, error) {
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

func (u *ContextUser) GetToken() string {
	return u.ctx.GetHeader("token")
}
