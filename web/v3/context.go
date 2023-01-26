package v3

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	PathParams map[string]string

	queryValues url.Values
}

type StringValue struct {
	val string
	err error
}

func newStringValue(val string, err error) StringValue {
	return StringValue{val: val, err: err}
}

func (c *Context) newDecoder() (error, *json.Decoder) {
	if c.Request.Body == nil {
		return errors.New("web: body为 nil"), nil
	}
	return nil, json.NewDecoder(c.Request.Body)
}

func (c *Context) SetCookie(ck *http.Cookie) {
	http.SetCookie(c.Response, ck)
}

// RespJSON 设置返回数据
func (c *Context) RespJSON(status int, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	c.Response.WriteHeader(status)
	//c.Response.Header().Set("Content-Type", "application/json")
	//c.Response.Header().Set("Content-Length", strconv.Itoa(len(data)))
	n, err := c.Response.Write(data)
	if n != len(data) {
		return errors.New("web:未写入全部数据")
	}
	return err
}

func (c *Context) RespJSONOK(val any) error {
	return c.RespJSON(http.StatusOK, val)
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

// BindJSONNumber 表示:数字就用Number来表示
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

// BindJSONDisallowUnknownFields 如果有一个未知的字段，就会保错
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

// FormValue 获取表单数据
func (c *Context) FormValue(key string) StringValue {
	err := c.Request.ParseForm()
	if err != nil {
		return newStringValue("", nil)
	}
	return newStringValue(c.Request.FormValue(key), nil)
}

// QueryValue 获取请求URL上的值
func (c *Context) QueryValue(key string) StringValue {
	if c.queryValues == nil {
		c.queryValues = c.Request.URL.Query()
	}
	vals, ok := c.queryValues[key]
	if !ok || len(vals) == 0 {
		return newStringValue("", errors.New("web: key 不存在"))
	}
	return newStringValue(vals[0], nil)
}

// PathValue 获取路由参数的值
func (c *Context) PathValue(key string) StringValue {
	val, ok := c.PathParams[key]
	if !ok {
		return newStringValue("", errors.New("web: key 不存在！"))
	}
	return newStringValue(val, nil)
}
