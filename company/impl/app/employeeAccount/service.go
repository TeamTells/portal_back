package employeeAccount

import (
	"context"
	"errors"
	"portal_back/authentication/api/internalapi"
	frontendapi "portal_back/company/api/frontend"
	"portal_back/core/network"
)

var EmployeeAlreadyExists = errors.New("employee with this email already exists in your company")
var EmployeeNotFound = errors.New("employee not found")

type Service interface {
	CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error
	GetEmployee(ctx context.Context, id int) (frontendapi.EmployeeWithConnections, error)
	DeleteEmployee(ctx context.Context, id int, requestInfo network.RequestInfo) error
	EditEmployee(ctx context.Context, id int, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error
	MoveEmployeesToDepartment(ctx context.Context, dto frontendapi.MoveEmployeeRequest) error
}

func NewService(repository Repository, authService internalapi.AuthRequestService) Service {
	return &service{repository: repository, authService: authService}
}

type service struct {
	repository  Repository
	authService internalapi.AuthRequestService
}

func (s *service) CreateEmployee(ctx context.Context, dto frontendapi.EmployeeRequest, requestInfo network.RequestInfo) error {
	//создаём юзера или берём уже созданного
	userId, err := s.createUser(ctx, dto.Email)
	if err != nil {
		return nil
	}

	//проверяем, состоит ли он в данной компании
	employee, err := s.repository.GetEmployeeByUserAndCompanyIds(ctx, userId, requestInfo.CompanyId)
	if err != nil {
		return err
	}
	//если уже состоит, кидаем ошибку
	if employee != nil {
		return EmployeeAlreadyExists
	}

	err = s.repository.CreateEmployee(ctx, dto, userId, requestInfo.CompanyId)
	if err != nil {
		return err
	}

	//после создания сотрудника добавляем его в департамент
	if dto.DepartmentId != nil {
		createdEmployee, err := s.repository.GetEmployeeByUserAndCompanyIds(ctx, userId, requestInfo.CompanyId)
		if err != nil {
			return err
		}
		err = s.repository.AddEmployeeToDepartment(ctx, *createdEmployee.Id, *dto.DepartmentId)
	}

	return nil
}

func (s *service) createUser(ctx context.Context, Email string) (int, error) {
	createUserErr := s.authService.CreateNewUser(ctx, Email)

	if createUserErr != nil && !errors.Is(createUserErr, internalapi.UserAlreadyExists) {
		return -1, createUserErr
	}

	userId, getUserErr := s.authService.GetUserIdByEmail(ctx, Email)

	if getUserErr != nil {
		return -1, getUserErr
	}

	return *userId, nil
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

func (s *service) MoveEmployeesToDepartment(ctx context.Context, dto frontendapi.MoveEmployeeRequest) error {
	//редактирование данных в таблице Employee_department
	for _, l := range *dto.Employees {
		err := s.repository.MoveEmployeeToDepartment(ctx, l.IdEmployee, l.IdDepartmentFrom, dto.IdDepartmentTo)
		if err != nil {
			return err
		}
	}
	return nil
}
