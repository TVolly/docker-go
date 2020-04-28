package repositories

type RepositoryStore interface {
	Community() CommunityRepository
}

func TestStore() RepositoryStore {
	return NewMemoryStore()
}
