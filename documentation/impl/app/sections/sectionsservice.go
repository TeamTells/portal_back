package sections

import (
	"portal_back/documentation/impl/domain"
)

func NewSectionService() SectionService {
	return &service{}
}

type SectionService interface {
	GetSections() ([]domain.Section, error)
}

type service struct {
}

func (s service) GetSections() ([]domain.Section, error) {
	sections := []domain.Section{
		{1, "First", ""},
		{2, "Second", ""},
	}
	return sections, nil
}
