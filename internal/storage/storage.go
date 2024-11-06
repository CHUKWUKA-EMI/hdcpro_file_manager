package storage

import (
	"fmt"
	"healthdecodepro_file_manager/internal/config"
	"healthdecodepro_file_manager/internal/model"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Storage is an interface that defines the methods that a storage implementation should implement
type Storage interface {
	GetUploadURL(location string, fileName string) (*model.PresignedUploadURLResponse, error)
	GetDeleteURL(location string, fileName string) (*model.PresignedDeleteURLResponse, error)
}

// S3Storage is a storage implementation that uses AWS S3
type S3Storage struct {
	App *config.AppConfig
}

const (
	// S3StorageType is the type of storage
	S3StorageType = "s3"
)

// GetUploadURL returns a pre-signed URL that can be used to upload a file to S3
func (s *S3Storage) GetUploadURL(location string, fileName string) (*model.PresignedUploadURLResponse, error) {
	s3Client := s3.New(s.App.AWS)

	request, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.App.StorageBucket),
		Key:    aws.String(location + "/" + fileName),
	})

	url, err := request.Presign(3 * time.Minute)
	if err != nil {
		return nil, err
	}

	response := &model.PresignedUploadURLResponse{
		UploadURL:   url,
		DownloadURL: fmt.Sprintf("%s/%s/%s", s.App.AssetsDistributionURL, location, fileName),
	}

	return response, nil
}

// GetDeleteURL returns a pre-signed URL that can be used to delete a file from S3
func (s *S3Storage) GetDeleteURL(location string, fileName string) (*model.PresignedDeleteURLResponse, error) {
	s3Client := s3.New(s.App.AWS)
	request, _ := s3Client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket: aws.String(s.App.StorageBucket),
		Key:    aws.String(location + "/" + fileName),
	})

	url, err := request.Presign(3 * time.Minute)
	if err != nil {
		return nil, err
	}

	return &model.PresignedDeleteURLResponse{
		DeleteURL: url,
	}, nil
}

// NewS3Storage creates a new S3Storage instance
func NewS3Storage(app *config.AppConfig) *S3Storage {
	return &S3Storage{App: app}
}

// NewStorage creates a new storage instance based on the provided type
func NewStorage(app *config.AppConfig, storageType string) Storage {
	switch storageType {
	case S3StorageType:
		return NewS3Storage(app)
	default:
		return nil
	}
}
