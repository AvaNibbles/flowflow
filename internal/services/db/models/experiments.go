package models

import "gorm.io/gorm"

type Namespace struct {
	gorm.Model
	Name string `gorm:"index:namespace_name_unique,unique"`
}

type Experiment struct {
	gorm.Model
	Name        string `gorm:"index:experiment_name_namespace_id_unique,unique"`
	NamespaceID int    `gorm:"index:experiment_name_namespace_id_unique,unique"`
	Namespace   Namespace
}
