package services

import (
	"errors"
)

const (
	id       = "1"
	name     = "zhang"
	password = "b5cf498b70a176efeacbc5b07d88e0da76a7f4cb"
)

type User struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

func NewUser(ID string, name string, password string) *User {
	return &User{ID: ID, Name: name, Password: password}
}

func GetInfoById(userId string) (*User, error) {
	if userId != id {
		return nil, errors.New("用户Id不存在！")
	}
	return NewUser(
		userId,
		name,
		password), nil
}
func GetInfo(userName string) (*User, error) {
	if name != userName {
		return nil, errors.New("用户名不存在！")
	}
	return NewUser(
		id,
		name,
		password), nil
}
