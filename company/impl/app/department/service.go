package department

import (
	"context"
	"errors"
	"portal_back/company/impl/domain"
	"portal_back/core/network"
)

type Service interface {
	GetDepartments(ctx context.Context, companyId int) ([]domain.DepartmentPreview, error)
	CreateDepartment(ctx context.Context, dto domain.DepartmentRequest, requestInfo network.RequestInfo) error
	GetDepartment(ctx context.Context, id int) (domain.Department, error)
	DeleteDepartment(ctx context.Context, id int, requestInfo network.RequestInfo) error
	EditDepartment(ctx context.Context, id int, dto domain.DepartmentRequest, requestInfo network.RequestInfo) error
	GetEmployees(ctx context.Context, companyId int) error
}

var DepartmentEmployeesNotFound = errors.New("employees in this department not found")

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type service struct {
	repository Repository
}

func (s *service) GetDepartments(ctx context.Context, companyId int) ([]domain.DepartmentPreview, error) {
	rootDepartments, err := s.repository.GetCompanyDepartments(ctx, companyId)
	if err != nil {
		return nil, err
	}
	var resultDeps []domain.DepartmentPreview
	for _, dep := range rootDepartments {
		count, _ := s.repository.GetCountOfEmployees(ctx, dep.Id)
		var arr []domain.DepartmentPreview
		depPreview := domain.DepartmentPreview{
			CountOfEmployees: count,
			Departments:      &arr, Id: dep.Id, Name: dep.Name,
		}
		resultDeps = append(resultDeps, depPreview)
		err := s.recursion(ctx, depPreview, func(d domain.DepartmentPreview) {
			arr = append(arr, d)
		})
		if err != nil {
			return nil, err
		}
	}
	return resultDeps, nil
}

func (s *service) recursion(ctx context.Context, department domain.DepartmentPreview, addToParentDepartment func(domain.DepartmentPreview)) error {
	childDepartments, err := s.repository.GetChildDepartments(ctx, department.Id)
	if err != nil {
		return err
	}
	for _, dep := range childDepartments {
		count, _ := s.repository.GetCountOfEmployees(ctx, dep.Id)
		var arr []domain.DepartmentPreview
		normDep := domain.DepartmentPreview{
			CountOfEmployees: count,
			Departments:      &arr, Id: dep.Id, Name: dep.Name,
		}
		err := s.recursion(ctx, normDep, func(d domain.DepartmentPreview) {
			arr = append(arr, d)
		})
		if err != nil {
			return err
		}
	}
	addToParentDepartment(department)
	return nil
}

func (s *service) CreateDepartment(ctx context.Context, dto domain.DepartmentRequest, requestInfo network.RequestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetDepartment(ctx context.Context, id int) (domain.Department, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteDepartment(ctx context.Context, id int, requestInfo network.RequestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) EditDepartment(ctx context.Context, id int, dto domain.DepartmentRequest, requestInfo network.RequestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetEmployees(ctx context.Context, companyId int) error {
	//TODO implement me
	panic("implement me")
}
