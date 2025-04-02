package service

import (
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
)

type CategoryService interface {
	PageQuery(dto dto.CategoryPageQueryDTO) (result.PageResult, error)
	Save(categoryDTO dto.CategoryDTO) error
	DeleteById(value int) error
	Update(categoryDTO dto.CategoryDTO) error
	StartOrStop(status int, id int) error
	List(kind int) ([]entity.Category, error)
}
