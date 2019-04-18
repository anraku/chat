package usecase

import "github.com/anraku/chat/domain/repository"

type UserUsecase interface {
	SaveUser(interface{}) error
}

type UserInteractor struct {
	ur repository.UserRepository
	mr repository.MessageRepository
}

func NewUserInteractor(ur repository.UserRepository, mr repository.MessageRepository) UserUsecase {
	return &UserInteractor{
		ur: ur,
		mr: mr,
	}
}

func (i *UserInteractor) SaveUser(data interface{}) error {
	err := i.ur.Create(data)
	return err
}
