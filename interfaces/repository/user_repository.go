package repository

import (
	"github.com/anraku/chat/entity"
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

func (r *UserRepository) Create(m *entity.Message) error {
	return r.DB.Create(m).Error
}
