package main

import (
	"context"
	"fmt"
	"oss-test/storage"
)

func main() {
	// 初始化参数
	storageType := "s3"                  // Change to "minio" for MinIO
	endpoint := "s3.amazonaws.com"       // AWS S3 endpoint
	accessKeyID := "your-access-key"     // AWS Access Key
	secretAccessKey := "your-secret-key" // AWS Secret Key
	useSSL := true                       // AWS S3 uses SSL
	bucketName := "your-bucket"          // Your S3 bucket name
	objectName := "your-object"          // Your S3 object name

	// 创建存储客户端
	client, err := storage.CreateStorageClient(storageType, endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 检查存储访问权限
	if exists, err := client.BucketExists(context.Background(), bucketName); err != nil {
		fmt.Println(err)
		return
	} else if !exists {
		fmt.Println("Bucket does not exist")
		return
	}

	if err := client.ObjectExists(context.Background(), bucketName, objectName); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Object %s exists in bucket %s\n", objectName, bucketName)
}
