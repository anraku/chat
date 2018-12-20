package usecase

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

func (i *UserInteractor) SaveUser(data interface{}) error {
	err := i.userRepository.Create(data)
	return err
}
