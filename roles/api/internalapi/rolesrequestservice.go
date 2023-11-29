package internalapi

import "portal_back/roles/api/internalapi/model"

type RolesRequestService interface {
	IsUserHasRole(accountId int, roleType model.RoleType) bool
}
