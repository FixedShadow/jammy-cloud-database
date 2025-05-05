package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// go操作 minio对象存储

// mysql内核数据包 mysql运行日志 mysql备份数据 以及mysql-monitor-agent新版本都将存在于对象存储中

const uploadFilePath = "./upload.log"
const uploadObjectName = "/go/upload.log"
const objectName = "/go/go1.20.3.linux-amd64.tar.gz"
const downloadFilePath = "./download.tar.gz"
const bucketName = "test-bucket"
const accessKeyId = "uwWGI0XA31EOlvSj0WPr"
const secretAccessKey = "3TaQkYpaJ5S7WalRnbyu8krLwXmVhABuCWYGdIC2"

func Download() {
	ctx := context.Background()
	endpoint := "192.168.2.88:9800"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	err = minioClient.FGetObject(ctx, bucketName, objectName, downloadFilePath, minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}
}

func Upload() {
	ctx := context.Background()
	endpoint := "192.168.2.88:9800"
	contentType := "application/octet-stream"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	_, err = minioClient.FPutObject(ctx, bucketName, uploadObjectName, uploadFilePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		panic(err)
	}
}
