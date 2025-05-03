package impl

import (
	"fmt"
	"sky-take-out/common/constant"
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/mapperParams"
	"time"
)

type CategoryServiceImpl struct{}

func (c *CategoryServiceImpl) PageQuery(dto dto.CategoryPageQueryDTO) (res result.PageResult, err error) {
	return mapperParams.CategoryMapper.PageQuery(dto)
}

func (c *CategoryServiceImpl) Save(categoryDTO dto.CategoryDTO) (err error) {
	category := entity.Category{
		Type:       functionParams.ToInt(categoryDTO.Type),
		Name:       categoryDTO.Name,
		Sort:       functionParams.ToInt(categoryDTO.Sort),
		Status:     constant.DISABLE,
		CreateUser: functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		UpdateUser: functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return mapperParams.CategoryMapper.Save(category)
}

func (c *CategoryServiceImpl) DeleteById(value int) error {
	count, err := mapperParams.DishMapper.CountByCategoryId(value)
	if err != nil || count > 0 {
		return fmt.Errorf(constant.CATEGORY_BE_RELATED_BY_DISH)
	}
	count, err = mapperParams.SetmealMapper.CountByCategoryId(value)
	if err != nil || count > 0 {
		return fmt.Errorf(constant.CATEGORY_BE_RELATED_BY_SETMEAL)
	}
	return mapperParams.CategoryMapper.DeleteById(value)
}

func (c *CategoryServiceImpl) Update(categoryDTO dto.CategoryDTO) error {
	category := entity.Category{
		ID:         int64(functionParams.ToInt(categoryDTO.Id)),
		Name:       categoryDTO.Name,
		Sort:       functionParams.ToInt(categoryDTO.Sort),
		Status:     -1,
		UpdateUser: functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		UpdateTime: time.Now(),
	}
	return mapperParams.CategoryMapper.Update(category)
}

func (c *CategoryServiceImpl) StartOrStop(status int, id int) error {
	category := entity.Category{
		ID:         int64(id),
		Sort:       -1,
		Status:     status,
		UpdateUser: -1,
	}
	return mapperParams.CategoryMapper.Update(category)
}

func (c *CategoryServiceImpl) List(kind int) ([]entity.Category, error) {
	return mapperParams.CategoryMapper.List(kind, constant.ENABLE)
}
