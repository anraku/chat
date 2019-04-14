package repository

type UserRepository interface {
	Create(interface{}) error
}
