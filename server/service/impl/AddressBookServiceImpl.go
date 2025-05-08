package impl

import (
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/mapperParams"
)

type AddressBookServiceImpl struct {
}

func (a *AddressBookServiceImpl) List(book entity.AddressBook) ([]entity.AddressBook, error) {
	return mapperParams.AddressBookMapper.List(book)
}

func (a *AddressBookServiceImpl) Save(book entity.AddressBook) error {
	book.UserID = int64(commonParams.Thread.Get()["userId"].(float64))
	book.IsDefault = 0
	return mapperParams.AddressBookMapper.Insert(book)
}

func (a *AddressBookServiceImpl) GetById(book entity.AddressBook) (entity.AddressBook, error) {
	books, err := mapperParams.AddressBookMapper.List(book)
	if err != nil {
		return entity.AddressBook{}, err
	}
	return books[0], err
}

func (a *AddressBookServiceImpl) Update(book entity.AddressBook) error {
	return mapperParams.AddressBookMapper.Update(book)
}

func (a *AddressBookServiceImpl) SetDefault(book entity.AddressBook) (err error) {
	book.IsDefault = 0
	book.UserID = int64(commonParams.Thread.Get()["userId"].(float64))
	err = mapperParams.AddressBookMapper.UpdateIsDefaultByUserId(book)
	if err != nil {
		return err
	}
	book.IsDefault = 1
	return mapperParams.AddressBookMapper.Update(book)

}

func (a *AddressBookServiceImpl) DeleteById(id int64) error {
	return mapperParams.AddressBookMapper.DeleteById(id)
}
