package domain

import frontendapi "portal_back/company/api/frontend"

type Department struct {
	Id               int
	Name             string
	ParentDepartment frontendapi.CommonEntity
	Company          frontendapi.CommonEntity
	Supervisor       frontendapi.CommonEntity
}
