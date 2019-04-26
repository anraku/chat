package usecase

import (
	"github.com/anraku/chat/domain/repository"
)

type UserUsecase interface {
	SaveUser(interface{}) error
}

type UserInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) UserUsecase {
	return &UserInteractor{
		userRepository: ur,
	}
}

func (i *UserInteractor) SaveUser(data interface{}) error {
	return i.userRepository.Create(data)
}
