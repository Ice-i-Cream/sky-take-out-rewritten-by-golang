package service

import (
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
)

type DishService interface {
	PageQuery(dto dto.DishPageQueryDTO) (result.PageResult, error)
	SaveWithFlavor(dishDTO dto.DishDTO) error
	DeleteBatch(list []string) error
	GetByIdWithFlavor(id int) (vo.DishVO, error)
	UpdateWithFlavor(dishDTO dto.DishDTO) error
	List(id int) ([]entity.Dish, error)
	StartOrStop(status int, id int) error
}
