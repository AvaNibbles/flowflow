package domain

import (
	"github.com/avanibbles/flowflow/internal/services/db/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MLModelService interface {
	Get(namespace string, name string) (*models.MLModel, error)
}

func newMLModelService(d *DomainFactory) MLModelService {
	return &mlModelService{
		db:     d.db,
		logger: d.logger.With(zap.String("component", "MLModelService")),
	}
}

type mlModelService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (s *mlModelService) Get(namespace string, name string) (*models.MLModel, error) {
	var ret models.MLModel
	s.db.Model(&ret).Joins("Namespace").First(&ret, &models.MLModel{Name: name, Namespace: models.Namespace{Name: namespace}})
	return &ret, nil
}
