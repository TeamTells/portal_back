package sections

import (
	"context"
	"portal_back/documentation/impl/domain"
)

type SectionService interface {
	GetSections(context context.Context, companyId int) ([]domain.Section, error)
	CreateSection(context context.Context, section domain.Section, organizationId int) error
	UpdateIsFavoriteSection(context context.Context, sectionId int, isFavorite bool) error
}

func NewSectionService(sectionRepository SectionRepository) SectionService {
	return &service{sectionRepository: sectionRepository}
}

type service struct {
	sectionRepository SectionRepository
}

func (service *service) CreateSection(context context.Context, section domain.Section, organizationId int) error {
	return service.sectionRepository.CreateSection(context, section, organizationId)
}

func (service *service) GetSections(context context.Context, companyId int) ([]domain.Section, error) {
	return service.sectionRepository.GetSections(context, companyId)
}

func (service *service) UpdateIsFavoriteSection(
	context context.Context,
	sectionId int,
	isFavorite bool,
) error {
	return service.sectionRepository.UpdateIsFavoriteSection(context, sectionId, isFavorite)
}
