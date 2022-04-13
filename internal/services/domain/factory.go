package domain

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Factory struct {
	db     *gorm.DB
	logger *zap.Logger
	wg     *sync.WaitGroup
}

func New(db *gorm.DB, logger *zap.Logger, wg *sync.WaitGroup) *Factory {
	return &Factory{
		db:     db,
		logger: logger,
		wg:     wg,
	}
}

func (d *Factory) NewMaintenanceService() MaintenanceService {
	return newMaintenanceService(d)
}

func (d *Factory) NewMLModelService() MLModelService {
	return newMLModelService(d)
}

func (d *Factory) NewNamespaceService() NamespaceService {
	return newNamespaceService(d)
}
