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
		INSERT INTO sections (title, thumbnail_url, company_id)
		VALUES ($1, $2, $3);
	`

	rows, error := repository.conn.Query(context, query, section.Title, section.ThumbnailUrl, organizationId)
	defer rows.Close()

	return error
}

func (repository sectionRepositoryImpl) GetSections(context context.Context, companyId int, userId int) ([]domain.Section, error) {
	query := `
		SELECT sections.id, sections.title, sections.thumbnail_url, COALESCE(user_sections_prefs.is_favorite, false) FROM sections
		LEFT JOIN user_sections_prefs ON sections.id=user_sections_prefs.section_id
		AND user_sections_prefs.user_id=$1
		WHERE company_id=$2
	`

	rows, err := repository.conn.Query(context, query, userId, companyId)
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

func (repository sectionRepositoryImpl) UpdateIsFavoriteSection(
	context context.Context,
	sectionId int,
	userId int,
	isFavorite bool,
) error {
	query := `
		INSERT INTO user_sections_prefs(user_id, section_id, is_favorite) 
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, section_id) DO UPDATE SET is_favorite=$3
	`

	rows, error := repository.conn.Query(context, query, userId, sectionId, isFavorite)
	defer rows.Close()

	return error
}
