package usecase

type UserUsecase interface {
	SaveUser(interface{}) error
}

type UserInteractor struct {
	userRepository    UserRepository
	messageRepository MessageRepository
}

func NewUserInteractor(ur UserRepository, mr MessageRepository) UserUsecase {
	return &UserInteractor{
		userRepository:    ur,
		messageRepository: mr,
	}
}

func (i *UserInteractor) SaveUser(data interface{}) error {
	err := i.userRepository.Create(data)
	return err
}
