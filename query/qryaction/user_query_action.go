package qryaction

import (
	"github.com/iniyusril/presence-api/entity"
	"github.com/iniyusril/presence-api/repository"
)

type UserQueryAction struct {
	UserRepository repository.UserRepository
}

func NewUserQueryAction(userRepository repository.UserRepository) UserQueryAction {
	return UserQueryAction{
		UserRepository: userRepository,
	}
}

func (s *UserQueryAction) FindAll() (*[]entity.User, error) {
	return s.UserRepository.FindAll()
}
