package service

import (
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/vo"
)

type SetmealService interface {
	PageQuery(dto dto.SetmealPageQueryDTO) (result.PageResult, error)
	SaveWithDish(setmealDTO dto.SetmealDTO) error
	DeleteBatch(setmeal []string) error
	GetByIdWithDish(param string) (vo.SetmealVO, error)
	Update(setmealDTO dto.SetmealDTO) error
	StartOrStop(setmealDTO dto.SetmealDTO) error
}
