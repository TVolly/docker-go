package repositories

import (
	"github.com/TVolly/goapi-addresses/internal/models"
)

type communityMemoryRepository struct {
	items map[int]*models.Community
}

func NewCommunityMemoryRepository() *communityMemoryRepository {
	return &communityMemoryRepository{
		items: make(map[int]*models.Community),
	}
}

func (r *communityMemoryRepository) Create(m *models.Community) error {
	m.ID = len(r.items) + 1

	r.items[m.ID] = m

	return nil
}

func (r *communityMemoryRepository) Find(id int) (*models.Community, error) {
	m, ok := r.items[id]

	if !ok {
		return nil, ErrRecordNotFound
	}

	return m, nil
}
