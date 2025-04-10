package repository

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SaveUser(user *entity.User) (entity.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return entity.User{}, err
	}
	return *user, nil
}
