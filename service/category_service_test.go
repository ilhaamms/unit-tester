package service

import (
	"testing"

	"github.com/ilhaamms/unit-testing/entity"
	"github.com/ilhaamms/unit-testing/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepositoryMock = repository.CategoryRepositoryMock{Mock: mock.Mock{}}  // manggil repo mock
var categoryService = CategoryService{categoryRepository: &categoryRepositoryMock} // inject repo mock di service

func TestCategoryService_NotFound(t *testing.T) {
	// program mock, jadi kita bebas disini mau return apa
	categoryRepositoryMock.Mock.On("FindById", "1").Return(nil) // On adalah ketika sebuah function

	result, err := categoryService.Get("1") // 1 ini adalah sama dengan parameter program mock yg udah di set

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestCategoryService_Success(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	categoryRepositoryMock.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
