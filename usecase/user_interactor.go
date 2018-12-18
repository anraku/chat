package usecase

import (
	"github.com/anraku/chat/entity"
	"github.com/anraku/chat/interfaces"
)

type UserInteractor struct {
	userRepository    interfaces.UserRepository
	messageRepository interfaces.MessageRepository
}

func NewUserInteractor(ur interfaces.UserRepository, mr interfaces.MessageRepository) *UserInteractor {
	return &UserInteractor{
		userRepository:    ur,
		messageRepository: mr,
	}
}

func (i *UserInteractor) StoreData(m *entity.Message) error {
	err := i.userRepository.Create(m)
	return err
}
