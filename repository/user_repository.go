package repository

import (
	"github.com/iniyusril/presence-api/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u *UserRepository) Save(user entity.User) (*entity.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepository) FindAll() (*[]entity.User, error) {
	var users []entity.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
