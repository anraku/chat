package repository

import (
	"github.com/anraku/chat/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(m *domain.Message) error {
	return r.DB.Create(m).Error
}
