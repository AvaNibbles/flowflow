package domain

import (
	"github.com/avanibbles/flowflow/internal/services/db/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NamespaceService interface {
	Get(name string) (*models.Namespace, error)
	Create(name string) (*models.Namespace, error)
	Delete(name string) error
}

func newNamespaceService(d *Factory) NamespaceService {
	return &namespaceService{
		db:     d.db,
		logger: d.logger.With(zap.String("component", "MLModelService")),
	}
}

type namespaceService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (n *namespaceService) Get(name string) (*models.Namespace, error) {
	var ret models.Namespace
	res := n.db.First(&ret, &models.Namespace{Name: name})
	if res.Error != nil {
		return nil, res.Error
	}

	return &ret, nil
}

func (n *namespaceService) Create(name string) (*models.Namespace, error) {
	newNs := models.Namespace{Name: name}
	res := n.db.Create(&newNs)
	if res.Error != nil {
		return nil, res.Error
	}

	return &newNs, nil
}

func (n *namespaceService) Delete(name string) error {
	res := n.db.Delete(&models.Namespace{}, &models.Namespace{Name: name})
	return res.Error
}
