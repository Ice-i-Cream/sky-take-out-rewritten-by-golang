package service

import "sky-take-out/pojo/entity"

type AddressBookService interface {
	List(book entity.AddressBook) ([]entity.AddressBook, error)
	Save(book entity.AddressBook) error
	GetById(book entity.AddressBook) (entity.AddressBook, error)
	Update(book entity.AddressBook) error
	SetDefault(book entity.AddressBook) error
	DeleteById(id int64) error
}
