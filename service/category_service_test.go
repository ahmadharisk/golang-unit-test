package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_FindById_NotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.FindById("1")
	assert.Nil(t, category)
	assert.Error(t, err)
}

func TestCategoryService_FindById_Success(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Makanan",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.FindById("2")
	assert.Nil(t, err)
	assert.Equal(t, category.Name, result.Name)
	assert.Equal(t, category.Id, result.Id)
}
