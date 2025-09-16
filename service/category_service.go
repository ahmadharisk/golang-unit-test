package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) FindById(id string) (*entity.Category, error) {
	categoryById := service.Repository.FindById(id)
	if categoryById == nil {
		return nil, errors.New("categoryById not found")
	}
	return categoryById, nil
}
