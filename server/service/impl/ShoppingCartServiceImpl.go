package impl

import (
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/mapperParams"
	"strconv"
	"time"
)

type ShoppingCartServiceImpl struct {
}

func (s *ShoppingCartServiceImpl) ShowShoppingCart() ([]entity.ShoppingCart, error) {
	userId := commonParams.Thread.Get()["userId"].(float64)
	shoppingCart := entity.ShoppingCart{
		UserID:    int64(userId),
		SetmealID: -1,
		DishID:    -1,
	}
	return mapperParams.ShoppingCartMapper.List(shoppingCart)

}

func (s *ShoppingCartServiceImpl) AddShoppingCart(dto dto.ShoppingCartDTO) error {
	userId := int64(commonParams.Thread.Get()["userId"].(float64))
	shoppingCart := entity.ShoppingCart{
		DishID:     dto.DishId,
		SetmealID:  dto.SetmealId,
		DishFlavor: dto.DishFlavor,
		UserID:     userId,
	}

	list, err := mapperParams.ShoppingCartMapper.List(shoppingCart)
	if err != nil {
		return err
	}
	if len(list) > 0 {
		cart := list[0]
		cart.Number = cart.Number + 1
		err := mapperParams.ShoppingCartMapper.UpdateNumberById(cart)
		if err != nil {
			return err
		}
	} else {
		if dto.DishId != 0 {
			dishId := dto.DishId
			dish, err := mapperParams.DishMapper.GetById(strconv.FormatInt(dishId, 10))
			if err != nil {
				return err
			}
			shoppingCart.Name = dish.Name
			shoppingCart.Image = dish.Image
			shoppingCart.Amount = dish.Price
		} else {
			setmealId := dto.SetmealId
			setmeal, err := mapperParams.SetmealMapper.GetById(int(setmealId))
			if err != nil {
				return err
			}
			shoppingCart.Name = setmeal.Name
			shoppingCart.Image = setmeal.Image
			shoppingCart.Amount = setmeal.Price
		}
		shoppingCart.Number = 1
		shoppingCart.CreateTime = time.Now()
		err := mapperParams.ShoppingCartMapper.Insert(shoppingCart)
		if err != nil {
			return err
		}
	}
	return nil
}
