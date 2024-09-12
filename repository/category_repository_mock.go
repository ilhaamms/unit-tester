package repository

import (
	"github.com/ilhaamms/unit-testing/entity"
	"github.com/stretchr/testify/mock"
)

// bikin mock karena si service pasti manggil repo
// nah kalau mau buat unit test yang diinject seperti itu maka wajib buat mock nya
type CategoryRepositoryMock struct {
	Mock mock.Mock // dari testify
}

// ini adalah method struct yang ada di interface repository
// ketika unit test manggil FindById, maka akan manggil si mock nya
func (categoryRepository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := categoryRepository.Mock.Called(id) // ini returnnya []Call
	if arguments.Get(0) == nil {                    // cek apakah data pertama itu nil
		return nil
	}

	category := arguments.Get(0).(entity.Category) // konversi menjadi category
	return &category
}
