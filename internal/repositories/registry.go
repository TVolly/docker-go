package repositories

type RepositoryRegistry interface {
	Community() CommunityRepository
}
