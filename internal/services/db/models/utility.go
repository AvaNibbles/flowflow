package models

import "gorm.io/gorm"

const (
	METADATA_KEY_LAST_PREPPED_VERSION = "last-prepped-version"
)

type ServiceMetadata struct {
	gorm.Model
	Key   string `gorm:"index:service_metadata_key_unique,unique"`
	Value string
}
