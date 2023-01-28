package v3

import (
	"bytes"
	"context"
	"html/template"
)

type TemplateEngine interface {
	//Render 渲染页面
	// tplName 模板名称
	// data 渲染数据
	Render(ctx context.Context, tplName string, data any) ([]byte, error)
}

type GoTemplateEngine struct {
	T *template.Template
}

func (g *GoTemplateEngine) Render(ctx context.Context, tplName string, data any) ([]byte, error) {
	bs := &bytes.Buffer{}
	err := g.T.ExecuteTemplate(bs, tplName, data)
	return bs.Bytes(), err
}
