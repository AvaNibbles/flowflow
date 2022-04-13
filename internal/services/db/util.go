package db

import (
	"errors"

	"github.com/avanibbles/flowflow/internal/util"
	"gorm.io/gorm"
)

func MapToStatusError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return util.MakeStatusErr(404, "record not found", err)
	}

	return util.MakeStatusErr(500, "database error", err)
}
