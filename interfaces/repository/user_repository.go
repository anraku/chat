package repository

import (
	"github.com/anraku/chat/usecase"
)

type UserSessionRepository struct{}

func NewUserSessionRepository() usecase.UserRepository {
	return &UserSessionRepository{}
}

func (r *UserSessionRepository) Create(data interface{}) error {
	// save user to session
	return nil
}
