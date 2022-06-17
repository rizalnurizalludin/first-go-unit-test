package service

import (
	"testing"

	"github.com/rizalnurizalludin/first-go-unit-test/entity"
	"github.com/rizalnurizalludin/first-go-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {

	//program mocking
	categoryRepository.Mock.On("FindbyId","1").Return(nil)

	category, err:=categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t,err)

}

func TestCategoryFound(t *testing.T) {
	category:=entity.Category{
		Id:"1",
		Name:"Laptop",
	}

	categoryRepository.Mock.On("FindbyId", "2").Return(category)
	result,err:=categoryService.Get("2")
	assert.Nil(t,err)
	assert.NotNil(t,result)
	assert.Equal(t,category.Id,result.Id)
	assert.Equal(t,category.Name,result.Name)

}
