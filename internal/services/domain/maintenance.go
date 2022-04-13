package domain

import (
	"errors"
	"sync"

	"github.com/avanibbles/flowflow/internal/services/db/models"
	"github.com/avanibbles/flowflow/internal/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MaintenanceService interface {
	PreServiceStart() error
	OnServiceStart() error
}

func newMaintenanceService(d *Factory) MaintenanceService {
	return &maintenanceService{
		db:     d.db,
		logger: d.logger.With(zap.String("component", "maintenance")),
		wg:     d.wg,
	}
}

type maintenanceService struct {
	db     *gorm.DB
	logger *zap.Logger
	wg     *sync.WaitGroup
}

func (m *maintenanceService) PreServiceStart() error {
	if err := models.Migrate(m.db); err != nil {
		return err
	}

	return nil
}

func (m *maintenanceService) OnServiceStart() error {
	go func() {
		if err := m.maintenanceStart(); err != nil {
			// we want to crash if this has issues
			panic(err)
		}

		util.LoopUntilCancel(m.wg, m.logger, 5000, m.maintenanceLoop)
	}()

	return nil
}

func (m *maintenanceService) maintenanceStart() error {
	if err := m.checkDataPrep(); err != nil {
		return err
	}

	return nil
}

func (m *maintenanceService) maintenanceLoop() error {
	return nil
}

func (m *maintenanceService) checkDataPrep() error {
	var lastPreppedVersion models.ServiceMetadata
	findRes := m.db.First(models.ServiceMetadata{Key: models.METADATA_KEY_LAST_PREPPED_VERSION}, &lastPreppedVersion)
	if errors.Is(findRes.Error, gorm.ErrRecordNotFound) {
		// do data prep
		m.logger.Info("service requires data prep")
	} else if lastPreppedVersion.Value == "vNext" {
		// dev data prep
	}

	// this is where we'll handle data prep / migration for future versions

	return nil
}
