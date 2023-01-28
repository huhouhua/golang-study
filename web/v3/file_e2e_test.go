//go:build e2e

package v3

import (
	"github.com/stretchr/testify/require"
	"html/template"
	"log"
	"mime/multipart"
	"path/filepath"
	"testing"
)

func TestUpload(t *testing.T) {

	tpl, err := template.ParseGlob("../../assets/tpls/*.gohtml")
	require.NoError(t, err)

	engine := &GoTemplateEngine{
		T: tpl,
	}
	h := NewHTTPServer(ServerWithTemplateEngine(engine))
	//获取上传文件页面
	h.Get("/upload", func(ctx *Context) {
		err := ctx.Render("upload.gohtml", nil)
		if err != nil {
			log.Println(err)
		}
	})

	//处理上传文件
	fu := FileUploader{
		FileField: "myfile",
		DstPathFunc: func(header *multipart.FileHeader) string {
			return filepath.Join("../../", "assets", "upload", header.Filename)
		},
	}
	h.Post("/upload", fu.Handler())
	h.Start(":8081")
}

func TestDownload(t *testing.T) {

	h := NewHTTPServer()

	fn := FileDownloader{
		Dir: filepath.Join("../../", "assets", "download"),
	}
	h.Get("/download", fn.Handler())
	h.Start(":8081")
}
