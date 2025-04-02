package impl

import (
	"fmt"
	"sky-take-out/common/constant"
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/mapperParams"
	"strconv"
	"time"
)

type DishServiceImpl struct{}

func (d *DishServiceImpl) PageQuery(dto dto.DishPageQueryDTO) (result.PageResult, error) {
	return mapperParams.DishMapper.PageQuery(dto)

}

func (d *DishServiceImpl) SaveWithFlavor(dishDTO dto.DishDTO) (err error) {

	commonParams.Tx, err = commonParams.Db.Begin()
	if err != nil {
		return err
	}

	dish := entity.Dish{
		ID:          int64(functionParams.ToInt(dishDTO.ID)),
		Name:        dishDTO.Name,
		CategoryID:  int64(functionParams.ToInt(dishDTO.CategoryID)),
		Price:       float64(functionParams.ToInt(dishDTO.Price)),
		Image:       dishDTO.Image,
		Description: dishDTO.Description,
		Status:      functionParams.ToInt(dishDTO.Status),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		UpdateUser:  functionParams.GetUser(commonParams.Thread.Get()["empId"]),
	}
	dishId, err := mapperParams.DishMapper.Insert(dish)
	if err != nil {
		_ = commonParams.Tx.Rollback()
		return err
	}

	list := dishDTO.Flavors
	if len(list) > 0 {
		err = mapperParams.DishFlavorMapper.InsertBatch(list, dishId)
		if err != nil {
			_ = commonParams.Tx.Rollback()
			return err
		}
	}

	return commonParams.Tx.Commit()
}

func (d *DishServiceImpl) DeleteBatch(list []string) (err error) {

	for _, id := range list {
		dish, err := mapperParams.DishMapper.GetById(id)
		if err != nil {
			return err
		}
		if dish.Status == constant.ENABLE {
			return fmt.Errorf(constant.DISH_ON_SALE)
		}
	}
	setmealIds, err := mapperParams.SetmealMapper.GetSetmealIdByDishIds(list)
	if err != nil {
		return err
	}
	if len(setmealIds) > 0 {
		return fmt.Errorf(constant.SETMEAL_ON_SALE)
	}

	if commonParams.Tx, err = commonParams.Db.Begin(); err != nil {
		return err
	}

	if err = mapperParams.DishMapper.DeleteByIds(list); err != nil {
		_ = commonParams.Tx.Rollback()
		return err
	}

	if err = mapperParams.DishFlavorMapper.DeleteByDishIds(list); err != nil {
		_ = commonParams.Tx.Rollback()
		return err
	}

	return commonParams.Tx.Commit()
}

func (d *DishServiceImpl) GetByIdWithFlavor(id int) (vo.DishVO, error) {
	dish, err := mapperParams.DishMapper.GetById(strconv.Itoa(id))
	if err != nil {
		return vo.DishVO{}, err
	}
	list, err := mapperParams.DishFlavorMapper.GetByDishId(id)
	if err != nil {
		return vo.DishVO{}, err
	}
	dishVO := vo.DishVO{
		ID:          dish.ID,
		Name:        dish.Name,
		CategoryID:  dish.CategoryID,
		Price:       dish.Price,
		Image:       dish.Image,
		Description: dish.Description,
		Status:      dish.Status,
		Flavors:     list,
		Time:        dish.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return dishVO, nil

}
