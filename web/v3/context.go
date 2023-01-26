package v3

import (
	"encoding/json"
	"errors"
	"net/http"
)

//var (
//	JSONUseNumber             = true
//	JSONDisallowUnknownFileds = true
//)

type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	PathParams map[string]string
}

func (c *Context) newDecoder() (error, *json.Decoder) {
	if c.Request.Body == nil {
		return errors.New("web: body为 nil"), nil
	}
	return nil, json.NewDecoder(c.Request.Body)
}

func (c *Context) BindJSON(val any) error {
	if val == "" {
		return errors.New("web: 输入为 nil")
	}
	err, decoder := c.newDecoder()
	if err != nil {
		return err
	}
	return decoder.Decode(val)
}

// 表示:数字就用Number来表示
// 否则默认是 float64
func (c *Context) BindJSONNumber(val any) error {
	if val == "" {
		return errors.New("web: 输入为 nil")
	}
	err, decoder := c.newDecoder()
	if err != nil {
		return err
	}
	decoder.UseNumber()

	return decoder.Decode(val)
}

// 如果有一个未知的字段，就会保错
// 比如说你 User只有Name和Email两个字段
// JSON 里面额外多了一个Age字段，那么就会报错
func (c *Context) BindJSONDisallowUnknownFields(val any) error {
	if val == "" {
		return errors.New("web: 输入为 nil")
	}
	err, decoder := c.newDecoder()
	if err != nil {
		return err
	}
	decoder.DisallowUnknownFields()

	return decoder.Decode(val)
}

func (c *Context) FormValue(key string) (string, error) {
	err := c.Request.ParseForm()
	if err != nil {
		return "", nil
	}
	return c.Request.FormValue(key), nil
}
