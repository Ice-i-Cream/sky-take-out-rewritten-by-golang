package service

import (
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
)

type UserService interface {
	WxLogin(dto dto.UserLoginDTO) (entity.User, error)
}
