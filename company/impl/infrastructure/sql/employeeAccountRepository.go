package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	frontendapi "portal_back/company/api/frontend"
	"portal_back/company/impl/app/employeeAccount"
)

func NewEmployeeAccountRepository(conn *pgx.Conn) employeeAccount.Repository {
	return &repository{conn: conn}
}

type repository struct {
	conn *pgx.Conn
}

func (r repository) CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest, userId int, companyId int) error {
	//TODO implement me
	panic("implement me")
	//INSERT
}

func (r repository) GetEmployee(ctx context.Context, id int) (frontendapi.EmployeeWithConnections, error) {
	//TODO implement me
	panic("implement me")
	//SELECT
}

func (r repository) GetEmployeeByUserAndCompanyIds(ctx context.Context, userId int, companyId int) (*frontendapi.EmployeeWithConnections, error) {
	//TODO implement me
	panic("implement me")
	//SELECT
}

func (r repository) DeleteEmployee(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
	//DELETE
}

func (r repository) EditEmployee(ctx context.Context, id int, dto frontendapi.EmployeeRequest) error {
	//TODO implement me
	panic("implement me")
	//UPDATE
}

func (r repository) MoveEmployeeToDepartment(ctx context.Context, employeeId int, departmentFromId *int, departmentToId int) error {
	//TODO implement me
	panic("implement me")
	//UPDATE
}

func (r repository) AddEmployeeToDepartment(ctx context.Context, employeeId int, departmentId int) error {
	//TODO implement me
	panic("implement me")
	//INSERT
}
