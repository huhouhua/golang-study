package v3

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	JSONUseNumber             = true
	JSONDisallowUnknownFileds = true
)

type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	PathParams map[string]string
}

func (c *Context) BindJSON(val any) error {
	if val == nil {
		return errors.New("web:输入为 nil")
	}
	if c.Request.Body == nil {
		return errors.New("web:body为 nil")
	}
	decoder := json.NewDecoder(c.Request.Body)
	//表示:数字就用Number来表示
	//否则默认是 float64
	if JSONUseNumber {
		decoder.UseNumber()
	}
	//如果有一个未知的字段，就会保错
	//比如说你 User只有Name和Email两个字段
	//JSON 里面额外多了一个Age字段，那么就会报错
	if JSONDisallowUnknownFileds {
		decoder.DisallowUnknownFields()
	}
	return decoder.Decode(val)
}
