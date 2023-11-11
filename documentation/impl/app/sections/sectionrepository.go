package sections

import (
	"context"
	"portal_back/documentation/impl/domain"
)

type SectionRepository interface {
	GetSections(context context.Context, companyId int) ([]domain.Section, error)
}
