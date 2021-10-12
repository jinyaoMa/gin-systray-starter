package models

import (
	"gorm.io/gorm"

	. "this/database"
)

type TestModel struct {
	*gorm.Model
}

func (t *TestModel) GetAll() (tests []*TestModel, err error) {
	result := DB.Find(&tests)
	err = result.Error
	return
}
