package datastore

import (
	"github.com/anraku/chat/domain/repository"
)

type UserSessionRepository struct{}

func NewUserSessionRepository() repository.UserRepository {
	return &UserSessionRepository{}
}

func (ur *UserSessionRepository) Create(data interface{}) error {
	// save user to session
	return nil
}
