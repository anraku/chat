package usecase

import "github.com/anraku/chat/domain/repository"

type UserInteractor struct {
	userRepository    repository.UserRepository
	messageRepository repository.MessageRepository
}

func NewUserInteractor(ur repository.UserRepository, mr repository.MessageRepository) *UserInteractor {
	return &UserInteractor{
		userRepository:    ur,
		messageRepository: mr,
	}
}

func (i *UserInteractor) SaveUser(data interface{}) error {
	err := i.userRepository.Create(data)
	return err
}
