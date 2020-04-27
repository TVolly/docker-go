package repositories

type memoryRepositoryRegistry struct {
	community *communityMemoryRepository
}

func NewMemoryRegistry() *memoryRepositoryRegistry {
	return &memoryRepositoryRegistry{}
}

func (s *memoryRepositoryRegistry) Community() CommunityRepository {
	if s.community == nil {
		s.community = NewCommunityMemoryRepository()
	}

	return s.community
}
