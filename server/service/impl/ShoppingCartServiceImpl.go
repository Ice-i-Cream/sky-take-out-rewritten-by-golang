package impl

import (
	"fmt"
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
		err = mapperParams.ShoppingCartMapper.UpdateNumberById(cart)
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
		err = mapperParams.ShoppingCartMapper.Insert(shoppingCart)
	}
	return err
}

func (s *ShoppingCartServiceImpl) CleanShoppingCart() (err error) {
	userId := commonParams.Thread.Get()["userId"].(float64)
	return mapperParams.ShoppingCartMapper.DeleteByUserId(int64(userId))
}

func (s *ShoppingCartServiceImpl) SubShoppingCart(dto dto.ShoppingCartDTO) error {
	userId := int64(commonParams.Thread.Get()["userId"].(float64))
	shoppingCart := entity.ShoppingCart{
		DishID:     dto.DishId,
		SetmealID:  dto.SetmealId,
		DishFlavor: dto.DishFlavor,
		UserID:     userId,
	}

	list, err := mapperParams.ShoppingCartMapper.List(shoppingCart)
	if err != nil || len(list) < 1 {
		return fmt.Errorf("shopping cart not exist")
	}
	cart := list[0]
	if cart.Number > 1 {
		cart.Number = cart.Number - 1
		err = mapperParams.ShoppingCartMapper.UpdateNumberById(cart)
	} else {
		err = mapperParams.ShoppingCartMapper.DeleteById(cart.ID)
	}
	return err
}
