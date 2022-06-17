package service

import (
	"errors"

	"github.com/rizalnurizalludin/first-go-unit-test/entity"
	"github.com/rizalnurizalludin/first-go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindbyId(id)
	if category != nil{
		return category, nil
	}else{
		return nil, errors.New("category not found")
	}
}
