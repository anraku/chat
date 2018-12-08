package main

type UserInteractor struct {
	userRepository    *UserRepository
	messageRepository *MessageRepository
}

func NewUserInteractor(ur *UserRepository, mr *MessageRepository) *UserInteractor {
	return &UserInteractor{
		userRepository:    ur,
		messageRepository: mr,
	}
}
