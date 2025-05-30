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
	"time"
)

type SetmealServiceImpl struct{}

func (s *SetmealServiceImpl) PageQuery(dto dto.SetmealPageQueryDTO) (result.PageResult, error) {

	return mapperParams.SetmealMapper.PageQuery(dto)
}

func (s *SetmealServiceImpl) SaveWithDish(setmealDTO dto.SetmealDTO) error {
	setmeal := entity.Setmeal{
		Id:          functionParams.ToInt(setmealDTO.Id),
		CategoryId:  setmealDTO.CategoryId,
		Name:        setmealDTO.Name,
		Price:       float64(functionParams.ToInt(setmealDTO.Price)),
		Status:      setmealDTO.Status,
		Description: setmealDTO.Description,
		Image:       setmealDTO.Image,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		UpdateUser:  functionParams.GetUser(commonParams.Thread.Get()["empId"]),
	}
	id, err := mapperParams.SetmealMapper.Insert(setmeal)
	if err != nil {
		return err
	}

	for index, _ := range setmealDTO.SetmealDishes {
		setmealDTO.SetmealDishes[index].SetmealID = int64(id)
	}

	return mapperParams.SetmealDishMapper.InsertBatch(setmealDTO.SetmealDishes)
}

func (s *SetmealServiceImpl) DeleteBatch(setmeal []string) (err error) {

	for _, i := range setmeal {
		id := functionParams.ToInt(i)
		setmeal, err := mapperParams.SetmealMapper.GetById(id)
		if err != nil {
			return err
		}
		if setmeal.Status == constant.ENABLE {
			return fmt.Errorf(constant.SETMEAL_ON_SALE)
		}
	}

	commonParams.Tx, _ = commonParams.Db.Begin()

	for _, i := range setmeal {
		id := functionParams.ToInt(i)
		err = mapperParams.SetmealMapper.DeleteById(id)
		if err != nil {
			functionParams.Rollback()
			return err
		}
		err = mapperParams.SetmealDishMapper.DeleteBySetmealId(id)
		if err != nil {
			functionParams.Rollback()
			return err
		}
	}

	return functionParams.Commit()
}

func (s *SetmealServiceImpl) GetByIdWithDish(i string) (setmealVO vo.SetmealVO, err error) {
	id := functionParams.ToInt(i)
	setmeal, err := mapperParams.SetmealMapper.GetById(id)
	if err != nil {
		return setmealVO, err
	}
	setmealVO = vo.SetmealVO{
		ID:          int64(setmeal.Id),
		CategoryID:  setmeal.CategoryId,
		Name:        setmeal.Name,
		Price:       setmeal.Price,
		Status:      setmeal.Status,
		Description: setmeal.Description,
		Image:       setmeal.Image,
		UpdateTime:  setmeal.UpdateTime,
		UTime:       setmeal.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	setmealVO.SetmealDishes, err = mapperParams.SetmealDishMapper.GetBySetmealId(id)
	return setmealVO, err
}

func (s *SetmealServiceImpl) Update(setmealDTO dto.SetmealDTO) (err error) {
	commonParams.Tx, _ = commonParams.Db.Begin()

	setmeal := entity.Setmeal{
		Id:          int(setmealDTO.Id),
		CategoryId:  setmealDTO.CategoryId,
		Name:        setmealDTO.Name,
		Price:       float64(functionParams.ToInt(setmealDTO.Price)),
		Status:      setmealDTO.Status,
		Description: setmealDTO.Description,
		Image:       setmealDTO.Image,
		UpdateTime:  time.Now(),
		UpdateUser:  functionParams.GetUser(commonParams.Thread.Get()["empId"]),
	}

	err = mapperParams.SetmealMapper.Update(setmeal)
	if err != nil {
		functionParams.Rollback()
		return err
	}
	err = mapperParams.SetmealDishMapper.DeleteBySetmealId(setmeal.Id)
	if err != nil {
		functionParams.Rollback()
		return err
	}
	for index, _ := range setmealDTO.SetmealDishes {
		setmealDTO.SetmealDishes[index].SetmealID = setmealDTO.Id
	}
	err = mapperParams.SetmealDishMapper.InsertBatch(setmealDTO.SetmealDishes)
	if err != nil {
		functionParams.Rollback()
		return err
	}

	return functionParams.Commit()
}

func (s *SetmealServiceImpl) StartOrStop(setmealDTO dto.SetmealDTO) error {
	setmeal := entity.Setmeal{
		Id:         functionParams.ToInt(setmealDTO.Id),
		CategoryId: -1,
		Price:      -1,
		Status:     setmealDTO.Status,
		UpdateUser: -1,
	}
	err := mapperParams.SetmealMapper.Update(setmeal)
	return err
}

func (s *SetmealServiceImpl) List(setmeal entity.Setmeal) ([]entity.Setmeal, error) {
	return mapperParams.SetmealMapper.List(setmeal)
}

func (s *SetmealServiceImpl) GetDishItemById(id int64) ([]vo.DishItemVO, error) {
	return mapperParams.SetmealMapper.GetDishItemBySetmealId(id)
}
