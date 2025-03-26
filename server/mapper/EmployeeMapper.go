package mapper

import (
	"database/sql"
	"log"
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
	selectSQL := "select * from employee "
	args := []interface{}{}
	res.Records = []interface{}{}
	if dto.Name != "" {
		selectSQL = selectSQL + " where name like concat('%', ? ,'%')"
		args = append(args, dto.Name)
	}
	args = append(args, dto.PageSize)
	args = append(args, (dto.Page-1)*dto.PageSize)
	selectSQL = selectSQL + " order by create_time desc limit ? offset ?"
	log.Println(selectSQL, dto.Name)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return res, err
	}
	res.Total = 0
	for rows.Next() {
		var employee entity.Employee
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Username, &employee.Password, &employee.Phone, &employee.Sex, &employee.IDNumber, &employee.Status, &employee.CreateTime, &employee.UpdateTime, &employee.CreateUser, &employee.UpdateUser)
		if err != nil {
			return res, err
		}
		res.Records = append(res.Records, employee)
		res.Total++
	}
	return res, err
}
