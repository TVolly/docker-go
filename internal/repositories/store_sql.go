package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type sqlStore struct {
	databaseURL string
	db          *sql.DB
	community   *communitySqlRepository
}

func NewSqlStore(databaseURL string) *sqlStore {
	return &sqlStore{
		databaseURL: databaseURL,
	}
}

func (s *sqlStore) Init() error {
	db, err := sql.Open("postgres", s.databaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *sqlStore) Community() CommunityRepository {
	if s.community == nil {
		s.community = NewCommunitySqlRepository(s.db)
	}

	return s.community
}
