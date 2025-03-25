package impl

import (
	"github.com/dgrijalva/jwt-go"
	"sky-take-out/common/utils"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/commonParams"
	"sky-take-out/server/mapper"
	"time"
)

type EmployeeServiceImpl struct{}

func (e *EmployeeServiceImpl) Login(dto dto.EmployeeLoginDTO) (data vo.EmployeeLoginVO, err error) {
	employee, err := mapper.GetByUsername(dto.Username)
	if err != nil {
		return data, err
	}
	claims := jwt.MapClaims{
		"claims": jwt.MapClaims{
			"empId": employee.ID,
		}}
	token, err := utils.GenToken(claims, commonParams.JwtProperties.AdminSecretKey)
	if err != nil {
		return data, err
	}
	employeeLoginVO := vo.EmployeeLoginVO{
		ID:       0,
		UserName: employee.Username,
		Name:     employee.Name,
		Token:    token,
	}
	commonParams.RedisDb.Set(commonParams.Ctx, token, token, time.Hour)
	return employeeLoginVO, nil
}
