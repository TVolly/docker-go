package repositories

import (
	"github.com/TVolly/goapi-addresses/internal/models"
)

type communityMemoryRepository struct {
	items []*models.Community
}

func NewCommunityMemoryRepository() *communityMemoryRepository {
	return &communityMemoryRepository{
		items: []*models.Community{},
	}
}

func (r *communityMemoryRepository) Create(m *models.Community) error {
	m.ID = len(r.items) + 1

	r.items = append(r.items, m)

	return nil
}

func (r *communityMemoryRepository) List() []*models.Community {
	return r.items
}

func (r *communityMemoryRepository) Find(id int) (*models.Community, error) {
	for _, model := range r.items {
		if model.ID == id {
			return model, nil
		}
	}

	return nil, ErrRecordNotFound
}
