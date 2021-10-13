package services

import (
	"this/database/models"
)

type TestService struct {
	msg *string `query:"msg"`
}

func (ts *TestService) GetAllTestModels() (testModels []*models.TestModel, err error) {
	testModels, err = (&models.TestModel{}).GetAll()
	return
}
