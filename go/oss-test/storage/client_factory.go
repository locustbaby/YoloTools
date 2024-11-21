package storage

import (
	"fmt"
)

// CreateStorageClient is a factory function to create the appropriate storage client based on the type.
func CreateStorageClient(storageType, endpoint, accessKeyID, secretAccessKey string, useSSL bool) (StorageClient, error) {
	switch storageType {
	case "s3":
		return NewS3Client(endpoint, accessKeyID, secretAccessKey, useSSL)
	case "minio":
		return NewMinIOClient(endpoint, accessKeyID, secretAccessKey, useSSL)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", storageType)
	}
}
