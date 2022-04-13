package domain

import (
	"github.com/avanibbles/flowflow/internal/services/db"
	"github.com/avanibbles/flowflow/internal/services/db/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MLModelService interface {
	Get(namespace string, name string) (*models.MLModel, error)
	Create(namespace string, name string) (*models.MLModel, error)
}

func newMLModelService(d *Factory) MLModelService {
	return &mlModelService{
		db:     d.db,
		logger: d.logger.With(zap.String("component", "MLModelService")),
		ns:     newNamespaceService(d),
	}
}

type mlModelService struct {
	db     *gorm.DB
	logger *zap.Logger
	ns     NamespaceService
}

func (s *mlModelService) Get(namespace string, name string) (*models.MLModel, error) {
	var ret models.MLModel
	result := s.db.Model(&ret).Joins("Namespace").First(&ret, &models.MLModel{Name: name, Namespace: models.Namespace{Name: namespace}})
	if result.Error != nil {
		return nil, db.MapToStatusError(result.Error)
	}

	return &ret, nil
}

func (s *mlModelService) Create(namespace string, name string) (*models.MLModel, error) {
	ns, err := s.ns.Get(namespace)
	if err != nil {
		return nil, err
	}

	newModel := models.MLModel{Name: name, NamespaceID: ns.ID}
	result := s.db.Create(&newModel)
	if result.Error != nil {
		return nil, db.MapToStatusError(result.Error)
	}

	return &newModel, nil
}
