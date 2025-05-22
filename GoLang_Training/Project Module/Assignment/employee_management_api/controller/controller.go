package controller

import (
	"database/sql"
	"employee_management_api/database"
	"employee_management_api/model"
	"errors"
)

func CreateEmployee(employee model.Employee) error {
	//databasepackage.DataBase connection
	_, err := database.DataBase.Exec("INSERT INTO employee (name,emailid,projectname,location,mobileno) VALUES ($1,$2,$3,$4,$5)",
		employee.Name, employee.EmailId, employee.ProjectName, employee.Location, employee.MobileNo)
	return err
}

func GetEmployees() ([]model.Employee, error) {
	rows, err := database.DataBase.Query("SELECT * FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.EmpId, &employee.Name, &employee.EmailId, &employee.ProjectName, &employee.Location, &employee.MobileNo); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func GetEmployeeById(empid int) (model.Employee, error) {
	var employee model.Employee
	row := database.DataBase.QueryRow("SELECT empid,name,emailid,projectname,location,mobileno FROM employee WHERE empid=$1", empid)
	err := row.Scan(&employee.EmpId, &employee.Name, &employee.EmailId, &employee.ProjectName, &employee.Location, &employee.MobileNo)
	if err == sql.ErrNoRows {
		return employee, errors.New("User Not Found")
	}
	return employee, err
}
