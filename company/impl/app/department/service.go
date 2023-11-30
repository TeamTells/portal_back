package department

import (
	"context"
	frontendapi "portal_back/company/api/frontend"
	"portal_back/core/network"
)

type Service interface {
	GetCompanyDepartments(ctx context.Context, companyId int) (*[]frontendapi.AllDepartmentsResponse, error)
	CreateNewDepartment(ctx context.Context, dto frontendapi.DepartmentRequest, requestInfo network.RequestInfo) error
	GetDepartment(ctx context.Context, id int) (frontendapi.Department, error)
	DeleteDepartment(ctx context.Context, id int, requestInfo network.RequestInfo) error
	EditDepartment(ctx context.Context, id int, dto frontendapi.DepartmentRequest, requestInfo network.RequestInfo) error
	GetCompanyDepartmentsWithEmployees(ctx context.Context, companyId int) error
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type service struct {
	repository Repository
}

func (s *service) GetCompanyDepartments(ctx context.Context, companyId int) (*[]frontendapi.AllDepartmentsResponse, error) {

	rootDepartments, err := s.repository.GetRootCompanyDepartments(ctx, companyId)
	if err != nil {
		return nil, err
	}
	var resultDeps []frontendapi.AllDepartmentsResponse

	for _, dep := range rootDepartments {
		count, _ := s.repository.GetCountOfDepartmentEmployees(ctx, dep.Id)
		var arr []frontendapi.AllDepartmentsResponse
		normDep := frontendapi.AllDepartmentsResponse{
			CountOfEmployees: &count,
			Departments:      &arr, Id: &dep.Id, Name: &dep.Name,
		}
		resultDeps = append(resultDeps, normDep)
		err := s.recursion(ctx, normDep, func(d frontendapi.AllDepartmentsResponse) {
			arr = append(arr, d)
		})
		if err != nil {
			return nil, err
		}
	}
	return &resultDeps, nil
}

func (s *service) recursion(ctx context.Context, department frontendapi.AllDepartmentsResponse, addToParentDepartment func(frontendapi.AllDepartmentsResponse)) error {
	childDepartments, err := s.repository.GetChildDepartments(ctx, *department.Id)
	if err != nil {
		return err
	}
	for _, dep := range childDepartments {
		count, _ := s.repository.GetCountOfDepartmentEmployees(ctx, dep.Id)
		var arr []frontendapi.AllDepartmentsResponse
		normDep := frontendapi.AllDepartmentsResponse{
			CountOfEmployees: &count,
			Departments:      &arr, Id: &dep.Id, Name: &dep.Name,
		}
		err := s.recursion(ctx, normDep, func(d frontendapi.AllDepartmentsResponse) {
			arr = append(arr, d)
		})
		if err != nil {
			return err
		}
	}
	addToParentDepartment(department)
	return nil
}

func (s *service) CreateNewDepartment(ctx context.Context, dto frontendapi.DepartmentRequest, requestInfo network.RequestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetDepartment(ctx context.Context, id int) (frontendapi.Department, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteDepartment(ctx context.Context, id int, requestInfo network.RequestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) EditDepartment(ctx context.Context, id int, dto frontendapi.DepartmentRequest, requestInfo network.RequestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetCompanyDepartmentsWithEmployees(ctx context.Context, companyId int) error {
	//TODO implement me
	panic("implement me")
}
