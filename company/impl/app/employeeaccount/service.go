package employeeaccount

import (
	"context"
	"errors"
	"portal_back/authentication/api/internalapi"
	"portal_back/company/impl/domain"
	"portal_back/core/network"
)

var EmployeeAlreadyExists = errors.New("employee with this email already exists in your company")
var EmployeeNotFound = errors.New("employee not found")

type Service interface {
	CreateEmployee(ctx context.Context, dto domain.EmployeeRequest, requestInfo network.RequestInfo) error
	GetEmployee(ctx context.Context, id int) (domain.EmployeeWithConnections, error)
	DeleteEmployee(ctx context.Context, id int, requestInfo network.RequestInfo) error
	EditEmployee(ctx context.Context, id int, dto domain.EmployeeRequest, requestInfo network.RequestInfo) error
	MoveEmployeesToDepartment(ctx context.Context, dto domain.MoveEmployeesRequest) error
}

func NewService(repository Repository, userService internalapi.UserRequestService) Service {
	return &service{repository: repository, userService: userService}
}

type service struct {
	repository  Repository
	userService internalapi.UserRequestService
}

func (s *service) CreateEmployee(ctx context.Context, dto domain.EmployeeRequest, requestInfo network.RequestInfo) error {
	userId, err := s.createUserOrGetExisting(dto.Email)
	if err != nil {
		return nil
	}

	_, err = s.repository.GetCompanyEmployee(ctx, userId, requestInfo.CompanyId)
	if !errors.Is(err, EmployeeNotFound) {
		return EmployeeAlreadyExists
	}

	err = s.repository.CreateEmployee(ctx, dto, userId, requestInfo.CompanyId)
	if err != nil {
		return err
	}

	if dto.DepartmentID != nil {
		createdEmployee, err := s.repository.GetCompanyEmployee(ctx, userId, requestInfo.CompanyId)
		if err != nil {
			return err
		}
		err = s.repository.AddEmployeeToDepartment(ctx, createdEmployee.Id, *dto.DepartmentID)
	}

	return nil
}

func (s *service) createUserOrGetExisting(Email string) (int, error) {
	createUserErr := s.userService.CreateNewUser(Email)

	if createUserErr != nil && !errors.Is(createUserErr, internalapi.UserAlreadyExists) {
		return -1, createUserErr
	}

	userId, getUserErr := s.userService.GetUserId(Email)

	if getUserErr != nil {
		return -1, getUserErr
	}

	return userId, nil
}

func (s *service) GetEmployee(ctx context.Context, id int) (domain.EmployeeWithConnections, error) {
	return s.repository.GetEmployee(ctx, id)
}

func (s *service) DeleteEmployee(ctx context.Context, id int, requestInfo network.RequestInfo) error {

	return nil
}

func (s *service) EditEmployee(ctx context.Context, id int, dto domain.EmployeeRequest, requestInfo network.RequestInfo) error {
	return nil
}

func (s *service) MoveEmployeesToDepartment(ctx context.Context, dto domain.MoveEmployeesRequest) error {
	for _, l := range dto.Employees {
		if l.DepartmentFromID != nil {
			err := s.repository.MoveEmployeeToDepartment(ctx, l.EmployeeID, *l.DepartmentFromID, dto.DepartmentToID)
			if err != nil {
				return err
			}
		} else {
			err := s.repository.AddEmployeeToDepartment(ctx, l.EmployeeID, dto.DepartmentToID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
