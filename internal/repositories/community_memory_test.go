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

func TestCommunityMemoryRepository_List(t *testing.T) {
	r := repositories.NewCommunityMemoryRepository()

	assert.Empty(t, r.List())

	r.Create(models.TestCommunity())
	r.Create(models.TestCommunity())

	assert.Len(t, r.List(), 2)
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

func TestCommunityMemoryRepository_Update(t *testing.T) {
	r := repositories.NewCommunityMemoryRepository()

	m1 := models.TestCommunity()
	m1.Name = "Some name"
	r.Create(m1)
	m1.Name = "Another name"
	assert.NoError(t, r.Update(m1))

	m2, _ := r.Find(m1.ID)
	assert.Equal(t, m1.Name, m2.Name)

	m3 := models.TestCommunity()
	m3.ID = 999
	assert.Error(t, r.Update(m3))
}
