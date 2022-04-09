package storage

import (
	"github.com/avanibbles/flowflow/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type s3StorageService struct {
	S3         *s3.S3
	BucketName string
	Logger     *zap.Logger
	Uploader   *s3manager.Uploader
}

func newS3StorageService(cfg *config.S3StorageConfig, logger *zap.Logger) (*s3StorageService, error) {
	awsConfig := &aws.Config{
		Credentials:      credentials.NewEnvCredentials(),
		Endpoint:         cfg.Endpoint,
		Region:           cfg.Region,
		DisableSSL:       cfg.DisableSSL,
		S3ForcePathStyle: cfg.ForcePathStyle,
	}

	session, _ := session2.NewSession(awsConfig)
	s3Client := s3.New(session)

	uploader := s3manager.NewUploader(session)

	return &s3StorageService{
		S3:         s3Client,
		BucketName: cfg.BucketName,
		Logger:     logger.With(zap.String("component", "S3StorageService")),
		Uploader:   uploader,
	}, nil
}

func (s *s3StorageService) Put(req PutRequest) (*PutResponse, error) {
	input := s3manager.UploadInput{Key: &req.Key, Body: req.Body, Bucket: &s.BucketName}
	resp, err := s.Uploader.Upload(&input)
	if err != nil {
		return nil, err
	}

	return &PutResponse{Location: resp.Location}, nil
}

func (s *s3StorageService) Get(req GetRequest) (*GetResponse, error) {
	input := s3.GetObjectInput{Bucket: &s.BucketName, Key: &req.Key}
	resp, err := s.S3.GetObject(&input)
	if err != nil {
		return nil, err
	}

	return &GetResponse{Body: resp.Body}, nil
}
