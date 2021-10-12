package models

import (
	. "this/database"

	"gorm.io/gorm"
)

type TestModel struct {
	*gorm.Model
}

func (t *TestModel) GetAll() (tests []*TestModel, err error) {
	result := DB.Find(&tests)
	err = result.Error
	return
}

func init() {
	DB.AutoMigrate(&TestModel{})
}
