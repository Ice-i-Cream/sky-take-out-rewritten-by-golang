package mapper

import (
	"database/sql"
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
)

func GetByUsername(username string) (empolyee entity.Employee, err error) {
	var employee entity.Employee
	var rows *sql.Rows
	selectSQL := "select * from sky_take_out.employee where username = ?"
	rows, err = commonParams.Db.Query(selectSQL, username)
	if err != nil {
		return empolyee, err
	}
	rows.Next()
	err = rows.Scan(&employee.ID, &employee.Username, &employee.Name, &employee.Password, &employee.Phone, &employee.Sex, &employee.IDNumber, &employee.Status, &employee.CreateTime, &employee.UpdateTime, &employee.CreateUser, &employee.UpdateUser)

	return employee, err
}

func Save(dto entity.Employee) (err error) {
	insertSQL := "insert into employee (name, username, password, phone, sex, id_number, status, create_time, update_time, create_user, update_user) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	_, err = commonParams.Db.Exec(insertSQL, dto.Name, dto.Username, dto.Password, dto.Phone, dto.Sex, dto.IDNumber, dto.Status, dto.CreateTime, dto.UpdateTime, dto.CreateUser, dto.UpdateUser)
	return err
}

func PageQuery(dto dto.EmployeePageQueryDTO) (res result.PageResult, err error) {
	return res, err
}
