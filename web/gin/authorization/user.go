package authorization

type UserInfo struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func NewUserInfo(ID string, name string, password string, token string) *UserInfo {
	return &UserInfo{ID: ID, Name: name, Password: password, Token: token}
}
