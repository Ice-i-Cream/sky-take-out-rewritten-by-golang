package service

import (
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
)

type ShoppingCartService interface {
	ShowShoppingCart() ([]entity.ShoppingCart, error)
	AddShoppingCart(dto dto.ShoppingCartDTO) error
	CleanShoppingCart() error
	SubShoppingCart(dto dto.ShoppingCartDTO) error
}
