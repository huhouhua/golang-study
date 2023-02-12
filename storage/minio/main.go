package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	endpoint        = "localhost:9000"
	accessKeyId     = "admin"
	secretAccessKey = "huhouhua"
	bucketName      = "rm.synyi.com"
	useSSL          = false
)

func main() {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(fmt.Sprintf("创建连接失败 %s", err.Error()))
	}
	//upload(minioClient)
	download(minioClient)

}

func upload(client *minio.Client) {
	objectName := "1.webp"
	path := filepath.Join("assets", "download", objectName)
	contentType := "application/webp"

	ctx := context.Background()
	defer ctx.Done()
	info, err := client.FPutObject(ctx, bucketName, objectName, path, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln("上传失败,", err)
		return
	}
	log.Printf("上传成功！name:%s  size:%d \n", objectName, info.Size)
}

func download(client *minio.Client) {
	objectName := "1.webp"
	ctx := context.Background()
	defer ctx.Done()
	object, err := client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})

	defer object.Close()

	if err != nil {
		log.Fatalln("下载失败！", err)
		return
	}
	path := filepath.Join("assets", "download", "new.webp")
	localFile, err := os.Create(path)
	if err != nil {
		log.Fatalln("创建文件失败", err)
		return
	}
	defer localFile.Close()

	if _, err := io.Copy(localFile, object); err != nil {
		fmt.Println("保存失败！", err)
	}
	log.Printf("下载成功！name:%s  \n", objectName)
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
