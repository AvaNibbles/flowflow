package models

import "gorm.io/gorm"

type MLModel struct {
	gorm.Model
	Name        string `gorm:"index:ml_model_name_namespace_id_unique,unique"`
	NamespaceID uint   `gorm:"index:ml_model_name_namespace_id_unique,unique"`
	Namespace   Namespace
}

type MLModelTag struct {
	gorm.Model
	TagName          string `gorm:"index:ml_model_tag_ml_model_id_unique,unique"`
	MLModelID        uint   `gorm:"index:ml_model_tag_ml_model_id_unique,unique"`
	MLModel          MLModel
	MLModelVersionID uint
	MLModelVersion   MLModelVersion
}

type MLModelVersion struct {
	gorm.Model
	MLModelID uint
	MLModel   MLModel
}

type MLModelVersionFile struct {
	gorm.Model
	MLModelVersionID uint
	MLModelVersion   MLModelVersion
}
