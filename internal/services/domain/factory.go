package domain

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DomainFactory struct {
	db     *gorm.DB
	logger *zap.Logger
	wg     *sync.WaitGroup
}

func New(db *gorm.DB, logger *zap.Logger, wg *sync.WaitGroup) *DomainFactory {
	return &DomainFactory{
		db:     db,
		logger: logger,
		wg:     wg,
	}
}

func (d *DomainFactory) NewMaintenanceService() MaintenanceService {
	return newMaintenanceService(d)
}

func (d *DomainFactory) NewMLModelService() MLModelService {
	return newMLModelService(d)
}
