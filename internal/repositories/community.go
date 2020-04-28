package repositories

import (
	"github.com/TVolly/goapi-addresses/internal/models"
)

type CommunityRepository interface {
	Create(m *models.Community) error
	List() []*models.Community
	Find(id int) (*models.Community, error)
}
