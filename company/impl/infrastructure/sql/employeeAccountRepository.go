package sql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"portal_back/company/impl/app/employeeaccount"
	"portal_back/company/impl/domain"
)

func NewEmployeeAccountRepository(conn *pgx.Conn) employeeaccount.Repository {
	return &employeeAccountRepository{conn: conn}
}

type employeeAccountRepository struct {
	conn *pgx.Conn
}

func (r employeeAccountRepository) CreateEmployee(ctx context.Context, dto domain.EmployeeRequest, userId int, companyId int) error {
	//TODO implement me
	panic("implement me")
}

func (r employeeAccountRepository) GetEmployee(ctx context.Context, id int) (domain.EmployeeWithConnections, error) {
	var employee domain.EmployeeWithConnections
	var rows pgx.Rows
	query := `
		SELECT employeeaccount.id, employeeaccount.firstname,
		       employeeaccount.secondname, employeeaccount.surname,
		       employeeaccount.dateofbirth, employeeaccount.telephonenumber,
		       employeeaccount.avatarurl, company.id, company.name, auth_user.email
		FROM employeeaccount
		LEFT JOIN company ON employeeaccount.companyid=company.id
		LEFT JOIN auth_user ON employeeaccount.userid=auth_user.id
        WHERE employeeaccount.id=$1
	`
	err := r.conn.QueryRow(ctx, query, id).Scan(&employee.Id, &employee.FirstName,
		&employee.SecondName, &employee.Surname,
		&employee.DateOfBirth, &employee.TelephoneNumber,
		&employee.Icon, &employee.Company.Id, &employee.Company.Name, &employee.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		return employee, employeeaccount.EmployeeNotFound
	}
	if err != nil {
		return employee, err
	}
	query = `
		SELECT department.id, department.name
		FROM department
		RIGHT JOIN employee_department ON employee_department.departmentid=department.id
		WHERE employee_department.accountid=$1
	`
	rows, err = r.conn.Query(ctx, query, id)
	for rows.Next() {
		var departmentInfo domain.DepartmentInfo
		err = rows.Scan(&departmentInfo.Id, &departmentInfo.Name)
		if err != nil {
			return employee, err
		}
		employee.Departments = append(employee.Departments, departmentInfo)
	}
	query = `
		SELECT role.id, role.title
		FROM role
		RIGHT JOIN employee_roles ON employee_roles.roleid=role.id
		WHERE employee_roles.accountid=$1
	`
	rows, err = r.conn.Query(ctx, query, id)
	for rows.Next() {
		var roleInfo domain.RoleInfo
		err = rows.Scan(&roleInfo.Id, &roleInfo.Name)
		if err != nil {
			return employee, err
		}
		employee.Roles = append(employee.Roles, roleInfo)
	}
	return employee, nil
}

func (r employeeAccountRepository) GetCompanyEmployee(ctx context.Context, userId int, companyId int) (domain.EmployeeWithConnections, error) {
	//TODO implement me
	panic("implement me")
}

func (r employeeAccountRepository) DeleteEmployee(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r employeeAccountRepository) EditEmployee(ctx context.Context, id int, dto domain.EmployeeRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r employeeAccountRepository) MoveEmployeeToDepartment(ctx context.Context, employeeId int, departmentFromId int, departmentToId int) error {
	//TODO implement me
	panic("implement me")
}

func (r employeeAccountRepository) AddEmployeeToDepartment(ctx context.Context, employeeId int, departmentId int) error {
	//TODO implement me
	panic("implement me")
}