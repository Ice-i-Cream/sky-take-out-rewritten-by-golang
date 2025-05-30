package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"strings"
)

type AddressBookMapper struct{}

func (a *AddressBookMapper) List(book entity.AddressBook) ([]entity.AddressBook, error) {
	selestSQL := "select * from address_book where true"
	args := []interface{}{}
	if book.ID != -1 {
		selestSQL = selestSQL + " and id=?"
		args = append(args, book.ID)
	}
	if book.UserID != -1 {
		selestSQL = selestSQL + " and user_id = ?"
		args = append(args, book.UserID)
	}
	if book.Phone != "" {
		selestSQL = selestSQL + " and phone = ?"
		args = append(args, book.Phone)
	}
	if book.IsDefault != -1 {
		selestSQL = selestSQL + " and is_default = ?"
		args = append(args, book.IsDefault)
	}
	log.Println(selestSQL, args)
	rows, err := commonParams.Db.Query(selestSQL, args...)
	if err != nil {
		log.Println(err)
	}
	var addressBooks []entity.AddressBook
	for rows.Next() {
		var book entity.AddressBook
		err := rows.Scan(&book.ID, &book.UserID, &book.Consignee, &book.Sex, &book.Phone, &book.ProvinceCode, &book.ProvinceName, &book.CityCode, &book.CityName, &book.DistrictCode, &book.DistrictName, &book.Detail, &book.Label, &book.IsDefault)
		if err != nil {
			return nil, err
		}
		addressBooks = append(addressBooks, book)
	}
	return addressBooks, nil
}

func (a *AddressBookMapper) Insert(book entity.AddressBook) error {
	insertSQL := "insert into address_book (user_id, consignee, sex, phone, province_code, province_name, city_code, city_name, district_code, district_name, detail, label, is_default) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	args := []interface{}{book.UserID, book.Consignee, book.Sex, book.Phone, book.ProvinceCode, book.ProvinceName, book.CityCode, book.CityName, book.DistrictCode, book.DistrictName, book.Detail, book.Label, book.IsDefault}
	log.Println(insertSQL, args)
	_, err := commonParams.Db.Exec(insertSQL, args...)
	return err
}

func (a *AddressBookMapper) Update(book entity.AddressBook) error {
	updateSQL := "update address_book set"
	args := []interface{}{}
	if book.Consignee != "" {
		updateSQL = updateSQL + " consignee = ?,"
		args = append(args, book.Consignee)
	}
	if book.Sex != "" {
		updateSQL = updateSQL + " sex = ?,"
		args = append(args, book.Sex)
	}
	if book.Phone != "" {
		updateSQL = updateSQL + " phone = ?,"
		args = append(args, book.Phone)
	}
	if book.Detail != "" {
		updateSQL = updateSQL + " detail = ?,"
		args = append(args, book.Detail)
	}
	if book.Label != "" {
		updateSQL = updateSQL + " label = ?,"
		args = append(args, book.Label)
	}
	if book.IsDefault != -1 {
		updateSQL = updateSQL + " is_default = ?,"
		args = append(args, book.IsDefault)
	}
	if book.DistrictCode != "" {
		updateSQL = updateSQL + " district_code = ?,"
		args = append(args, book.DistrictCode)
	}
	if book.DistrictName != "" {
		updateSQL = updateSQL + " district_name = ?,"
		args = append(args, book.DistrictName)
	}
	if book.ProvinceCode != "" {
		updateSQL = updateSQL + " province_code = ?,"
		args = append(args, book.ProvinceCode)
	}
	if book.ProvinceName != "" {
		updateSQL = updateSQL + " province_name = ?,"
		args = append(args, book.ProvinceName)
	}
	if book.CityCode != "" {
		updateSQL = updateSQL + " city_code = ?,"
		args = append(args, book.CityCode)
	}
	if book.CityName != "" {
		updateSQL = updateSQL + " city_name = ?,"
		args = append(args, book.CityName)
	}
	updateSQL = strings.TrimSuffix(updateSQL, ",") + " where id = ?"
	args = append(args, book.ID)
	log.Println(updateSQL, args)
	_, err := commonParams.Db.Exec(updateSQL, args...)
	return err
}

func (a *AddressBookMapper) UpdateIsDefaultByUserId(book entity.AddressBook) error {
	updateIsDefaultSQL := "update address_book set is_default = ? where user_id = ?"
	args := []interface{}{book.IsDefault, book.UserID}
	log.Println(updateIsDefaultSQL, args)
	_, err := commonParams.Db.Exec(updateIsDefaultSQL, args...)
	return err
}

func (a *AddressBookMapper) DeleteById(id int64) error {
	deleteSQL := "delete from address_book where id = ?"
	args := []interface{}{id}
	log.Println(deleteSQL, args)
	_, err := commonParams.Db.Exec(deleteSQL, args...)
	return err
}

func (a *AddressBookMapper) GetById(id int64) (book entity.AddressBook, err error) {
	selectSQL := "select * from address_book where id = ?"
	args := []interface{}{id}
	log.Println(selectSQL, args)
	row := commonParams.Db.QueryRow(selectSQL, args...)
	err = row.Scan(&book.ID, &book.UserID, &book.Consignee, &book.Sex, &book.Phone, &book.ProvinceCode, &book.ProvinceName, &book.CityCode, &book.CityName, &book.DistrictCode, &book.DistrictName, &book.Detail, &book.Label, &book.IsDefault)
	return book, err
}
