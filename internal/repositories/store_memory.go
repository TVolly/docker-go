package repositories

type memoryStore struct {
	community *communityMemoryRepository
}

func NewMemoryStore() *memoryStore {
	return &memoryStore{}
}

func (s *memoryStore) Community() CommunityRepository {
	if s.community == nil {
		s.community = NewCommunityMemoryRepository()
	}

	return s.community
}
