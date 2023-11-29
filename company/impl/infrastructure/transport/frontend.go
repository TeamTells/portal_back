package transport

import (
	"net/http"
	"portal_back/company/api/frontend"
	"portal_back/company/impl/app/employeeAccount"
	"portal_back/roles/api/internalapi"
)

func NewServer() frontendapi.ServerInterface {
	return &frontendServer{}
}

type frontendServer struct {
	employeeAccountService employeeAccount.Service
	rolesService           internalapi.RolesRequestService
}

func (f frontendServer) GetCompanyDepartments(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f frontendServer) CreateNewDepartment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}

func (f frontendServer) GetDepartment(w http.ResponseWriter, r *http.Request, departmentId int) {
	//TODO implement me
	panic("implement me")
}

func (f frontendServer) DeleteDepartment(w http.ResponseWriter, r *http.Request, departmentId int) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}

func (f frontendServer) EditDepartment(w http.ResponseWriter, r *http.Request, departmentId int) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}

func (f frontendServer) GetCompanyDepartmentsWithEmployees(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f frontendServer) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}

func (f frontendServer) MoveEmployeesToDepartment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}

func (f frontendServer) GetEmployee(w http.ResponseWriter, r *http.Request, employeeId int) {
	//TODO implement me
	panic("implement me")
}

func (f frontendServer) DeleteEmployee(w http.ResponseWriter, r *http.Request, employeeId int) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}

func (f frontendServer) EditEmployee(w http.ResponseWriter, r *http.Request, employeeId int) {
	//TODO implement me
	panic("implement me")
	//проверка на права (обращение в модуль ролей)
}
