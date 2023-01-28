package v3

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FileUploader struct {
	FileField   string
	DstPathFunc func(header *multipart.FileHeader) string
}

type FileDownloader struct {
	Dir string
}

func (d FileDownloader) Handler() HandlerFunc {
	return func(ctx *Context) {
		str := ctx.QueryValue("file")
		if str.Err != nil {
			ctx.ResponseStatusCode = http.StatusBadRequest
			ctx.ResponseData = []byte("找不到目标文件")
			return
		}
		dst := filepath.Join(d.Dir, str.Val)
		fn := filepath.Base(dst)

		header := ctx.Response.Header()
		header.Set("Content-Disposition", "attachment;filename="+fn)
		header.Set("Content-Description", "File Transfer")
		header.Set("Content-Type", "application/octet-stream")
		header.Set("Content-Transfer-Encoding", "binary")
		header.Set("Expires", "0")
		header.Set("Cache-Control", "must-revalidate")
		header.Set("Pragma", "public")

		http.ServeFile(ctx.Response, ctx.Request, dst)

	}
}

type FileUploaderOption func(uploader *FileUploader)

func NewFileUploader(opts ...FileUploaderOption) *FileUploader {
	res := &FileUploader{
		FileField: "file",
		DstPathFunc: func(header *multipart.FileHeader) string {
			return filepath.Join("../../", "upload", uuid.New().String())
		},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func (u FileUploader) Handler() HandlerFunc {
	return func(ctx *Context) {
		//上传文件的逻辑

		file, fileHeader, err := ctx.Request.FormFile(u.FileField)
		if err != nil {
			ctx.ResponseStatusCode = http.StatusInternalServerError
			ctx.ResponseData = []byte("上传失败" + err.Error())
			return
		}
		defer file.Close()
		//生成文件保存目录
		dst := u.DstPathFunc(fileHeader)

		//O_WRONLY 写入数据
		//O_TRUNC 如果文件不存在，清空数据
		//O_CREATE 创建一个新的
		dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o666)

		defer dstFile.Close()
		//写入数据
		_, err = io.CopyBuffer(dstFile, file, nil)
		if err != nil {
			ctx.ResponseStatusCode = http.StatusInternalServerError
			ctx.ResponseData = []byte("写入数据失败！" + err.Error())
		}
		ctx.ResponseStatusCode = http.StatusOK
		ctx.ResponseData = []byte("上传成功！")

	}
}
