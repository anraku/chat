package usecase

import (
	"github.com/anraku/chat/entity"
)

type UserInteractor struct {
	userRepository    UserRepository
	messageRepository MessageRepository
}

func NewUserInteractor(ur UserRepository, mr MessageRepository) *UserInteractor {
	return &UserInteractor{
		userRepository:    ur,
		messageRepository: mr,
	}
}

func (i *UserInteractor) StoreData(m *entity.Message) error {
	err := i.userRepository.Create(m)
	return err
}
