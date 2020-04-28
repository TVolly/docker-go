package repositories

type RepositoryStore interface {
	Init() error
	Community() CommunityRepository
}

func TestStore() RepositoryStore {
	return NewMemoryStore()
}
