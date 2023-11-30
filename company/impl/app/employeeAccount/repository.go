package employeeAccount

import (
	"context"
	frontendapi "portal_back/company/api/frontend"
)

type Repository interface {
	CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest, userId int, companyId int) error
	GetEmployee(ctx context.Context, id int) (frontendapi.EmployeeWithConnections, error)
	GetEmployeeByUserAndCompanyIds(ctx context.Context, userId int, companyId int) (*frontendapi.EmployeeWithConnections, error)
	DeleteEmployee(ctx context.Context, id int) error
	EditEmployee(ctx context.Context, id int, dto frontendapi.EmployeeRequest) error
	MoveEmployeeToDepartment(ctx context.Context, employeeId int, departmentFromId *int, departmentToId int) error
	AddEmployeeToDepartment(ctx context.Context, employeeId int, departmentId int) error
}
