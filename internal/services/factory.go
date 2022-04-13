package services

import (
	"sync"

	"github.com/avanibbles/flowflow/internal/services/db"
	"github.com/avanibbles/flowflow/internal/services/domain"
	"github.com/avanibbles/flowflow/internal/services/storage"
	"github.com/avanibbles/flowflow/pkg/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DependencyFactory interface {
	GetDomain() *domain.Factory
	GetStorage() storage.Service
	GetLogger(component string) *zap.Logger
	GetDb() *gorm.DB
	GetConfig() *config.Config
}

type DependencyConfig struct {
	Logger *zap.Logger
	Config *config.Config
	Wg     *sync.WaitGroup
}

func NewDependencyFactory(config *DependencyConfig) (DependencyFactory, error) {
	s, err := storage.NewStorageService(config.Config.Storage, config.Logger)
	if err != nil {
		return nil, err
	}

	dbInstance, err := db.New(config.Config.Db, config.Logger)
	if err != nil {
		return nil, err
	}

	domainFactory := domain.New(dbInstance, config.Logger, config.Wg)

	return &dependencyFactory{logger: config.Logger, storage: s, db: dbInstance, domain: domainFactory, config: config.Config}, nil
}

type dependencyFactory struct {
	logger *zap.Logger
	config *config.Config

	storage storage.Service
	db      *gorm.DB
	domain  *domain.Factory
}

func (d *dependencyFactory) GetStorage() storage.Service {
	return d.storage
}

func (d *dependencyFactory) GetLogger(component string) *zap.Logger {
	return d.logger.With(zap.String("component", component))
}

func (d *dependencyFactory) GetDomain() *domain.Factory {
	return d.domain
}

func (d *dependencyFactory) GetDb() *gorm.DB {
	return d.db
}

func (d *dependencyFactory) GetConfig() *config.Config {
	return d.config
}
