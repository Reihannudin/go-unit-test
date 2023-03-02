package repository

import "unitTest/model"

type CategoryRepository interface {
	FindById(id string) *model.Category
}
