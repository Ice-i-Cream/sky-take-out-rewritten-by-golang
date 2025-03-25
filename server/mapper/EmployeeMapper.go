package mapper

import (
	"database/sql"
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
