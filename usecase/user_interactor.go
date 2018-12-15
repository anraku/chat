package usecase

import (
	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/repository"
)

type UserInteractor struct {
	userRepository    *repository.UserRepository
	messageRepository *repository.MessageRepository
}

func NewUserInteractor(ur *repository.UserRepository, mr *repository.MessageRepository) *UserInteractor {
	return &UserInteractor{
		userRepository:    ur,
		messageRepository: mr,
	}
}

func (i *UserInteractor) storeData(m *domain.Message) error {
	err := i.userRepository.Create(m)
	return err
}
