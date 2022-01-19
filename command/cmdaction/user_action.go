package cmdaction

import (
	"github.com/iniyusril/presence-api/entity"
	"github.com/iniyusril/presence-api/repository"
)

type UserCommandAction struct {
	UserRepository repository.UserRepository
}

func NewUserCommandAction(userRepository repository.UserRepository) UserCommandAction {
	return UserCommandAction{
		UserRepository: userRepository,
	}
}

func (s *UserCommandAction) Save(user entity.User) (*entity.User, error) {
	res, err := s.UserRepository.Save(user)
	return res, err
}

func (s *UserCommandAction) FindAll() (*[]entity.User, error) {
	res, err := s.UserRepository.FindAll()
	return res, err
}
