package repository

import "github.com/ilhaamms/unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
