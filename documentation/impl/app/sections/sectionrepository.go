package sections

import (
	"context"
	"portal_back/documentation/impl/domain"
)

type SectionRepository interface {
	GetSections(context context.Context, companyId int) ([]domain.Section, error)
	CreateSection(context context.Context, section domain.Section, organizationId int) error
	UpdateIsFavoriteSection(context context.Context, sectionId int, isFavorite bool) error
}
