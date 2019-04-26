package usecase

import (
	"github.com/anraku/chat/domain/repository"
)

type UserUsecase interface {
	SaveUser(interface{}) error
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SaveUser(data interface{}) error {
	return uu.ur.Create(data)
}
