package services

import "this/server/models"

type TestService struct {
	// msg *string `json:"msg"`
}

func (ts *TestService) GetAllTestModels() (testModels []*models.TestModel, err error) {
	testModels, err = (&models.TestModel{}).GetAll()
	return
}
