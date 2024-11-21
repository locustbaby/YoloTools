package storage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// StorageClient defines the interface for cloud storage clients.
type StorageClient interface {
	BucketExists(ctx context.Context, bucketName string) (bool, error)
	ObjectExists(ctx context.Context, bucketName, objectName string) error
}

// S3Client implements the StorageClient interface for AWS S3.
type S3Client struct {
	client *minio.Client
}

// NewS3Client creates a new S3Client.
func NewS3Client(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*S3Client, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("Error initializing S3 client: %v", err)
	}
	return &S3Client{client: client}, nil
}

// BucketExists checks if the specified bucket exists in S3.
func (s *S3Client) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return s.client.BucketExists(ctx, bucketName)
}

// ObjectExists checks if the specified object exists in S3.
func (s *S3Client) ObjectExists(ctx context.Context, bucketName, objectName string) error {
	_, err := s.client.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	return err
}

// MinIOClient implements the StorageClient interface for MinIO.
type MinIOClient struct {
	client *minio.Client
}

// NewMinIOClient creates a new MinIOClient.
func NewMinIOClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*MinIOClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("Error initializing MinIO client: %v", err)
	}
	return &MinIOClient{client: client}, nil
}

// BucketExists checks if the specified bucket exists in MinIO.
func (m *MinIOClient) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return m.client.BucketExists(ctx, bucketName)
}

// ObjectExists checks if the specified object exists in MinIO.
func (m *MinIOClient) ObjectExists(ctx context.Context, bucketName, objectName string) error {
	_, err := m.client.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	return err
}
