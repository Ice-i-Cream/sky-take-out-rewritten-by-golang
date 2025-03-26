package impl

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"sky-take-out/common/constant"
	"sky-take-out/common/result"
	"sky-take-out/common/utils"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
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
	if employee.Password != utils.Md5(dto.Password) {
		return data, errors.New("password error")
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

func (e *EmployeeServiceImpl) Save(dto dto.EmployeeDTO) (err error) {
	employee := entity.Employee{
		Name:       dto.Name,
		Username:   dto.Username,
		Password:   utils.Md5(constant.DEFAULT_PASSWORD),
		Phone:      dto.Phone,
		Sex:        dto.Sex,
		IDNumber:   dto.IDNumber,
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateUser: int64(int(commonParams.Thread.Get()["empId"].(float64))),
		UpdateUser: int64(int(commonParams.Thread.Get()["empId"].(float64))),
	}
	return mapper.Save(employee)
}

func (e *EmployeeServiceImpl) PageQuery(dto dto.EmployeePageQueryDTO) (res result.PageResult, err error) {
	res, err = mapper.PageQuery(dto)
	return res, nil
}
