package service

import (
	"errors"

	"github.com/ilhaamms/unit-testing/entity"
	"github.com/ilhaamms/unit-testing/repository"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func (categoryService *CategoryService) Get(id string) (*entity.Category, error) {
	category := categoryService.categoryRepository.FindById(id)
	if category == nil {
		return nil, errors.New("not found")
	}

	return category, nil

}
