package teststore

import (
	"github.com/yerassylbolatov/http-rest-api/internal/app/model"
	"github.com/yerassylbolatov/http-rest-api/internal/app/store"
)
import _ "github.com/lib/pq"

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
