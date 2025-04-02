package mapper

import (
	"database/sql"
	"fmt"
	"log"
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"strings"
	"time"
)

type EmployeeMapper struct{}

func (e *EmployeeMapper) GetByUsername(username string) (empolyee entity.Employee, err error) {
	var employee entity.Employee
	var rows *sql.Rows
	selectSQL := "select * from sky_take_out.employee where username = ?"
	log.Println(selectSQL, username)
	rows, err = commonParams.Db.Query(selectSQL, username)
	if err != nil {
		return empolyee, err
	}
	rows.Next()
	err = rows.Scan(&employee.ID, &employee.Name, &employee.Username, &employee.Password, &employee.Phone, &employee.Sex, &employee.IDNumber, &employee.Status, &employee.CreateTime, &employee.UpdateTime, &employee.CreateUser, &employee.UpdateUser)

	return employee, err
}

func (e *EmployeeMapper) GetById(id int) (empolyee entity.Employee, err error) {
	var employee entity.Employee
	var rows *sql.Rows
	selectSQL := "select * from sky_take_out.employee where id = ?"
	log.Println(selectSQL, id)
	rows, err = commonParams.Db.Query(selectSQL, id)
	if err != nil {
		return empolyee, err
	}
	rows.Next()
	err = rows.Scan(&employee.ID, &employee.Name, &employee.Username, &employee.Password, &employee.Phone, &employee.Sex, &employee.IDNumber, &employee.Status, &employee.CreateTime, &employee.UpdateTime, &employee.CreateUser, &employee.UpdateUser)
	employee.Time = employee.UpdateTime.Format("2006-01-02 15:04:05")
	return employee, err
}

func (e *EmployeeMapper) Save(dto entity.Employee) (err error) {
	insertSQL := "insert into employee (name, username, password, phone, sex, id_number, status, create_time, update_time, create_user, update_user) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	log.Println(insertSQL)
	_, err = commonParams.Db.Exec(insertSQL, dto.Name, dto.Username, dto.Password, dto.Phone, dto.Sex, dto.IDNumber, dto.Status, dto.CreateTime, dto.UpdateTime, dto.CreateUser, dto.UpdateUser)
	return err
}

func (e *EmployeeMapper) PageQuery(dto dto.EmployeePageQueryDTO) (res result.PageResult, err error) {
	selectSQL := "select * from employee"
	var args []interface{}
	res.Records = []interface{}{}
	if dto.Name != "" {
		selectSQL = selectSQL + " where name like concat('%', ? ,'%')"
		args = append(args, dto.Name)
	}
	args = append(args, dto.PageSize)
	args = append(args, (dto.Page-1)*dto.PageSize)
	selectSQL = selectSQL + " order by create_time desc limit ? offset ?"
	log.Println(selectSQL, args)
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
		employee.Time = employee.UpdateTime.Format("2006-01-02 15:04:05")
		res.Records = append(res.Records, employee)
		res.Total++
	}
	return res, err
}

func (e *EmployeeMapper) Update(emp entity.Employee) (err error) {
	updateSQL := "update employee set" + ""
	var args []interface{}
	var updates []string
	if emp.Name != "" {
		updates = append(updates, " name = ?")
		args = append(args, emp.Name)
	}
	if emp.Username != "" {
		updates = append(updates, " username = ?")
		args = append(args, emp.Username)
	}
	if emp.Password != "" {
		updates = append(updates, " password = ?")
		args = append(args, emp.Password)
	}
	if emp.Phone != "" {
		updates = append(updates, " phone = ?")
		args = append(args, emp.Phone)
	}
	if emp.Sex != "" {
		updates = append(updates, " sex = ?")
		args = append(args, emp.Sex)
	}
	if emp.IDNumber != "" {
		updates = append(updates, " id_number = ?")
		args = append(args, emp.IDNumber)
	}
	if emp.UpdateTime != *new(time.Time) {
		updates = append(updates, " update_time = ?")
		args = append(args, emp.UpdateTime)
	}
	if emp.UpdateUser != -1 {
		updates = append(updates, " update_user = ?")
		args = append(args, emp.UpdateUser)
	}
	if emp.Status != -1 {
		updates = append(updates, " status = ?")
		args = append(args, emp.Status)
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	updateSQL += strings.Join(updates, ",")
	updateSQL = updateSQL + " WHERE id = ?"
	args = append(args, emp.ID)

	log.Println(updateSQL, args)
	_, err = commonParams.Db.Exec(updateSQL, args...)
	return err
}
