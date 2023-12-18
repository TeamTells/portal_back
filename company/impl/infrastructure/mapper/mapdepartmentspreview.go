package mapper

import (
	frontendapi "portal_back/company/api/frontend"
	"portal_back/company/impl/domain"
)

func MapDepartmentsPreview(departments []domain.DepartmentPreview) []frontendapi.Departments {
	var result []frontendapi.Departments
	var childDepartments []frontendapi.Departments
	for _, d := range departments {
		if d.Departments != nil && len(*d.Departments) > 0 {
			childDepartments = MapDepartmentsPreview(*d.Departments)
		}
		result = append(result, frontendapi.Departments{CountOfEmployees: d.CountOfEmployees, Departments: &childDepartments, Id: d.Id, Name: d.Name})
	}
	return result
}
