package rdbms

import (
	"github.com/jinzhu/gorm"
)

func GetDBError(db *gorm.DB) error {

	if db.Error == nil {
		return nil
	}

	return db.Error
}