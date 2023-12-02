package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"portal_back/company/impl/app/employeeAccount"
	"portal_back/company/impl/domain"
)

func NewEmployeeAccountRepository(conn *pgx.Conn) employeeAccount.Repository {
	return &repository{conn: conn}
}

type repository struct {
	conn *pgx.Conn
}

func (r repository) CreateEmployee(ctx context.Context, dto domain.EmployeeRequest, userId int, companyId int) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetEmployee(ctx context.Context, id int) (domain.EmployeeWithConnections, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetEmployeeByUserAndCompanyIds(ctx context.Context, userId int, companyId int) (*domain.EmployeeWithConnections, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteEmployee(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) EditEmployee(ctx context.Context, id int, dto domain.EmployeeRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) MoveEmployeeToDepartment(ctx context.Context, employeeId int, departmentFromId *int, departmentToId int) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) AddEmployeeToDepartment(ctx context.Context, employeeId int, departmentId int) error {
	//TODO implement me
	panic("implement me")
}
