package repositories

type RepositoryStore interface {
	Community() CommunityRepository
}
