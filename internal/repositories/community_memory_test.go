package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TVolly/goapi-addresses/internal/models"
	"github.com/TVolly/goapi-addresses/internal/repositories"
)

func TestCommunityMemoryRepository_Create(t *testing.T) {
	r := repositories.NewCommunityMemoryRepository()

	m := models.TestCommunity()

	assert.Equal(t, 0, m.ID)
	assert.NoError(t, r.Create(m))
	assert.NotNil(t, m.ID)
}

func TestCommunityMemoryRepository_Find(t *testing.T) {
	r := repositories.NewCommunityMemoryRepository()

	m1 := models.TestCommunity()
	r.Create(m1)

	m2, err := r.Find(m1.ID + 1)
	assert.Error(t, err)
	assert.Nil(t, m2)

	m3, err := r.Find(m1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, m3)
}
