package tests

import (
	"github.com/avanibbles/flowflow/internal/services"
	"github.com/avanibbles/flowflow/internal/services/db/models"
)

func dbDataPrep(df services.DependencyFactory) error {
	db := df.GetDb()

	var defaultNs models.Namespace

	db.FirstOrCreate(&defaultNs, &models.Namespace{Name: "default"})
	db.FirstOrCreate(&models.MLModel{}, &models.MLModel{Name: "test", NamespaceID: defaultNs.ID})

	return nil
}
