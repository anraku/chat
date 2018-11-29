package main

type UserInteractor struct {
	Repository *UserRepository
}

func NewUserInteractor(r *UserRepository) *UserInteractor {
	return &UserInteractor{
		Repository: r,
	}
}
