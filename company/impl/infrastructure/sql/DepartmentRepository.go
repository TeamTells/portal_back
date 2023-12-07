package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"portal_back/company/impl/app/department"
	"portal_back/company/impl/domain"
)

func NewDepartmentRepository(conn *pgx.Conn) department.Repository {
	return &repository{conn: conn}
}

type departmentRepository struct {
	conn *pgx.Conn
}

func (r repository) GetDepartment(ctx context.Context, id int) (*domain.Department, error) {

}

func (r repository) GetDepartmentEmployees(ctx context.Context, departmentId int) ([]domain.Employee, error) {

}

func (r repository) GetChildDepartments(ctx context.Context, id int) ([]domain.Department, error) {
	query := `
		SELECT department.id, department.name, department.parentdepartmentid, parentDepartment.name, 
		       department.supervisorid, employeeaccount.firstname
		FROM department
		LEFT JOIN employeeaccount ON department.supervisorid = employeeaccount.id
		LEFT JOIN department AS parentDepartment ON department.parentdepartmentid = department.id
		WHERE department.parentdepartmentid = $1
	`
	var childDepartments []domain.Department
	rows, err := r.conn.Query(ctx, query)
	if err == pgx.ErrNoRows {
		return childDepartments, department.DepartmentEmployeesNotFound
	} else if err != nil {
		return childDepartments, err
	}
	defer rows.Close()

	for rows.Next() {
		var childDepartment domain.Department
		err := rows.Scan(&childDepartment.Id, &childDepartment.Name, &childDepartment.ParentDepartment.Id,
			&childDepartment.ParentDepartment.Name, &childDepartment.Supervisor.Id, &childDepartment.Supervisor.Name)
		if err != nil {
			return childDepartments, err
		}
		childDepartments = append(childDepartments, childDepartment)
	}

	return childDepartments, nil
}

func (r repository) GetCountOfDepartmentEmployees(ctx context.Context, departmentId int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM employee_department
		WHERE id = $1
	`
	var countOfDepartmentEmployees int
	err := r.conn.QueryRow(ctx, query, departmentId).Scan(&countOfDepartmentEmployees)
	if err == pgx.ErrNoRows {
		return 0, department.DepartmentEmployeesNotFound
	} else if err != nil {
		return 0, err
	}

	return countOfDepartmentEmployees, nil
}

func (r repository) GetRootCompanyDepartments(ctx context.Context, companyId int) ([]domain.Department, error) {
	query := `
		SELECT department.id, department.name, department.supervisorid, employeeaccount.firstname
		FROM department
		LEFT JOIN employeeaccount ON department.supervisorid = employeeaccount.id
		WHERE department.companyid = $1 AND department.parentdepartmentid = NULL
	`
	var rootCompanyDepartments []domain.Department
	rows, err := r.conn.Query(ctx, query)
	if err == pgx.ErrNoRows {
		return rootCompanyDepartments, department.DepartmentEmployeesNotFound
	} else if err != nil {
		return rootCompanyDepartments, err
	}
	defer rows.Close()

	for rows.Next() {
		var rootCompanyDepartment domain.Department
		rootCompanyDepartment.ParentDepartment.Id = 0
		rootCompanyDepartment.ParentDepartment.Name = ``
		err := rows.Scan(&rootCompanyDepartment.Id, &rootCompanyDepartment.Name, &rootCompanyDepartment.Supervisor.Id,
			&rootCompanyDepartment.Supervisor.Name)
		if err != nil {
			return rootCompanyDepartments, err
		}
		rootCompanyDepartments = append(rootCompanyDepartments, rootCompanyDepartment)
	}

	return rootCompanyDepartments, nil
}
