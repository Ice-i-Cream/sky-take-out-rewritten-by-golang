package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type AddressBookController struct{}

func (a *AddressBookController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		addressBook := entity.AddressBook{
			UserID:    int64(commonParams.Thread.Get()["userId"].(float64)),
			IsDefault: -1,
			ID:        -1,
		}
		return serviceParams.AddressBookService.List(addressBook)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (a *AddressBookController) Save(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		addressBook := entity.AddressBook{}
		err = ctx.ShouldBind(&addressBook)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.AddressBookService.Save(addressBook)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (a *AddressBookController) GetById(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		addressBook := entity.AddressBook{
			ID:        int64(functionParams.ToInt(ctx.Param("id"))),
			IsDefault: -1,
			UserID:    -1,
		}

		return serviceParams.AddressBookService.GetById(addressBook)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (a *AddressBookController) Update(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		addressBook := entity.AddressBook{}
		err = ctx.ShouldBind(&addressBook)
		addressBook.IsDefault = -1
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.AddressBookService.Update(addressBook)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (a *AddressBookController) SetDefault(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		addressBook := entity.AddressBook{}
		err = ctx.ShouldBind(&addressBook)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.AddressBookService.SetDefault(addressBook)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)

}

func (a *AddressBookController) Delete(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := int64(functionParams.ToInt(ctx.Query("id")))
		return nil, serviceParams.AddressBookService.DeleteById(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (a *AddressBookController) GetDefault(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		addressBook := entity.AddressBook{
			UserID:    int64(commonParams.Thread.Get()["userId"].(float64)),
			IsDefault: 1,
			ID:        -1,
		}
		books, err := serviceParams.AddressBookService.List(addressBook)
		if len(books) == 0 {
			return nil, fmt.Errorf("没有查询到默认地址")
		}
		return books[0], err

	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
