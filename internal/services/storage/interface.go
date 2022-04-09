package storage

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/avanibbles/flowflow/pkg/config"
	"go.uber.org/zap"
)

func NewStorageService(cfg *config.StorageConfig, logger *zap.Logger) (Service, error) {
	switch strings.ToLower(cfg.Type) {
	case "s3":
		return newS3StorageService(cfg.S3, logger)
	default:
		return nil, errors.New(fmt.Sprintf("unknown storage type \"%s\"", cfg.Type))
	}
}

type Service interface {
	Put(req PutRequest) (*PutResponse, error)
	Get(req GetRequest) (*GetResponse, error)
}

type PutRequest struct {
	Key  string
	Body io.Reader
}

type PutResponse struct {
	Location string
}

type GetRequest struct {
	Key string
}

type GetResponse struct {
	Body io.Reader
}
