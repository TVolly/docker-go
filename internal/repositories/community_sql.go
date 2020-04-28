package repositories

import (
	"database/sql"

	"github.com/TVolly/goapi-addresses/internal/models"
)

type communitySqlRepository struct {
	db *sql.DB
}

func NewCommunitySqlRepository(db *sql.DB) *communitySqlRepository {
	return &communitySqlRepository{
		db: db,
	}
}

func (r *communitySqlRepository) Create(m *models.Community) error {
	if err := m.Validate(); err != nil {
		return err
	}

	return r.db.QueryRow(
		"INSERT INTO communities (name) VALUES ($1) RETURNING id",
		m.Name,
	).Scan(&m.ID)
}

func (r *communitySqlRepository) List() []*models.Community {
	return []*models.Community{}
}

func (r *communitySqlRepository) Find(id int) (*models.Community, error) {
	return nil, ErrRecordNotFound
}

func (r *communitySqlRepository) Update(m *models.Community) error {
	return ErrRecordNotFound
}
