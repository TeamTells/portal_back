package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/domain"
)

func NewSectionRepository(conn *pgx.Conn) sections.SectionRepository {
	return sectionRepositoryImpl{conn: conn}
}

type sectionRepositoryImpl struct {
	conn *pgx.Conn
}

func (repository sectionRepositoryImpl) CreateSection(
	context context.Context,
	section domain.Section,
	organizationId int,
) error {
	query := `
		INSERT INTO sections (title, thumbnail_url, is_favorite, company_id)
		VALUES ($1, $2, $3, $4);
	`

	rows, error := repository.conn.Query(context, query, section.Title, section.ThumbnailUrl, false, organizationId)
	defer rows.Close()

	return error
}

func (repository sectionRepositoryImpl) GetSections(context context.Context, companyId int) ([]domain.Section, error) {
	query := `
		SELECT id, title, thumbnail_url, is_favorite FROM sections 
        WHERE company_id=$1
	`

	rows, err := repository.conn.Query(context, query, companyId)
	defer rows.Close()

	var sections []domain.Section
	for rows.Next() {
		var section domain.Section
		rows.Scan(&section.Id, &section.Title, &section.ThumbnailUrl, &section.IsFavorite)
		sections = append(sections, section)
	}

	if err == pgx.ErrNoRows {
		return []domain.Section{}, nil
	} else if err != nil {
		return nil, err
	}

	return sections, nil
}
