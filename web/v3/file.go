package v3

import (
	"github.com/google/uuid"
	lru "github.com/hashicorp/golang-lru"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

type StaticResourceHandlerOption func(handler *StaticResourceHandler)

type StaticResourceHandler struct {
	dir                     string
	cache                   *lru.Cache
	extensionContextTypeMap map[string]string
	//大文件缓存
	maxSize int
}

func NewStaticResourceHandler(dir string, opts ...StaticResourceHandlerOption) (*StaticResourceHandler, error) {
	c, err := lru.New(100 * 1024 * 1024)
	if err != nil {
		return nil, err
	}
	res := &StaticResourceHandler{
		dir:   dir,
		cache: c,
		//10 MB
		maxSize: 1024 * 1024 * 10,
		extensionContextTypeMap: map[string]string{
			"jpeg": "image/jpeg",
			"jpe":  "image/jpe",
			"jpg":  "image/jpg",
			"png":  "image/png",
			"pdf":  "image/pdf",
			"webp": "image/webp",
		},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func StaticWithMaxFileSize(maxSize int) StaticResourceHandlerOption {
	return func(handler *StaticResourceHandler) {
		handler.maxSize = maxSize
	}
}
func StaticWithCahe(maxSize int) StaticResourceHandlerOption {
	return func(handler *StaticResourceHandler) {
		handler.maxSize = maxSize
	}
}
func (s *StaticResourceHandler) Handle(ctx *Context) {

	str := ctx.PathValue("file")
	if str.Err != nil {
		ctx.ResponseStatusCode = http.StatusBadRequest
		ctx.ResponseData = []byte("请求路径不对！")
		return
	}
	dst := filepath.Join(s.dir, str.Val)
	ext := filepath.Ext(dst)
	header := ctx.Response.Header()
	if data, ok := s.cache.Get(str.Val); ok {
		header.Set("Content-Type", s.extensionContextTypeMap[ext[1:]])
		header.Set("Content-Length", strconv.Itoa(len(data.([]byte))))
		ctx.ResponseData = data.([]byte)
		ctx.ResponseStatusCode = http.StatusOK
		return
	}
	data, err := os.ReadFile(dst)
	if err != nil {
		ctx.ResponseStatusCode = http.StatusInternalServerError
		ctx.ResponseData = []byte("服务器错误！")
		return
	}
	if len(data) <= s.maxSize {
		s.cache.Add(str.Val, data)
	}
	header.Set("Content-Type", s.extensionContextTypeMap[ext[1:]])
	header.Set("Content-Length", strconv.Itoa(len(data)))
	ctx.ResponseData = data
	ctx.ResponseStatusCode = http.StatusOK

}
