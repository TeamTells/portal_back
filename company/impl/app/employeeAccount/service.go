package employeeAccount

import (
	"context"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/api/internalapi/model"
	frontendapi "portal_back/company/api/frontend"
	"portal_back/core/network"
)

type Service interface {
	CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error
	GetEmployee(ctx context.Context, id int) (frontendapi.EmployeeWithConnections, error)
	DeleteEmployee(ctx context.Context, id int, requestInfo network.RequestInfo) error
	EditEmployee(ctx context.Context, id int, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error
	MoveEmployeesToDepartment(ctx context.Context, dto frontendapi.MoveEmployeeRequest, requestInfo network.RequestInfo) error
}

func NewService(repository Repository, authService internalapi.AuthRequestService) Service {
	return &service{repository: repository, authService: authService}
}

type service struct {
	repository  Repository
	authService internalapi.AuthRequestService
}

func (s *service) CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error {
	//создать user через модуль авторизации
	s.authService.CreateNewUser(ctx, model.CreateUserRequest{})
	//создать employeeAccount
	//создать связь с департаментом

	return nil
}

func (s *service) GetEmployee(ctx context.Context, id int) (frontendapi.EmployeeWithConnections, error) {
	// один запрос в репо, где подтягивается employee со всеми нужными связями (User, Company, Department)
	return frontendapi.EmployeeWithConnections{}, nil
}

func (s *service) DeleteEmployee(ctx context.Context, id int, requestInfo network.RequestInfo) error {

	//удаление
	return nil
}

func (s *service) EditEmployee(ctx context.Context, id int, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error {
	//редактирование данных в таблице EmployeeAccount
	//редактирование данных в таблице User
	return nil
}

func (s *service) MoveEmployeesToDepartment(ctx context.Context, dto frontendapi.MoveEmployeeRequest, requestInfo network.RequestInfo) error {
	//редактирование данных в таблице Employee_department
	return nil
}
