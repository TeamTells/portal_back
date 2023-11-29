package employeeAccount

import (
	"context"
	frontendapi "portal_back/company/api/frontend"
)

type Repository interface {
	CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest) error
	GetEmployee(ctx context.Context, id int) (frontendapi.EmployeeWithConnections, error)
	DeleteEmployee(ctx context.Context, id int) error
	EditEmployee(ctx context.Context, id int, dto frontendapi.EmployeeRequest) error
	MoveEmployeesToDepartment(ctx context.Context, dto frontendapi.MoveEmployeeRequest) error
	AddEmployeeToDepartment(ctx context.Context, employeeId int, departmentId int) error
}
