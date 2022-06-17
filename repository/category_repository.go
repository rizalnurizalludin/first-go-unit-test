package repository

import "github.com/rizalnurizalludin/first-go-unit-test/entity"

type CategoryRepository interface {
	FindbyId(id string) *entity.Category
}
