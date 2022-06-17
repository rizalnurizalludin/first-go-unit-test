package repository

import (
	"github.com/rizalnurizalludin/first-go-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindbyId(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) != nil{
		category := arguments.Get(0).(entity.Category)
		return &category
	}else{
		return nil
	}

}
