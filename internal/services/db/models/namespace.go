package models

import "gorm.io/gorm"

type Namespace struct {
	gorm.Model
	Name string `gorm:"index:namespace_name_unique,unique"`
}
