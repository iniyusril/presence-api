package cmdservice

import (
	"github.com/iniyusril/presence-api/command/cmdaction"
)

type UserService struct {
	UserCommandAction cmdaction.UserCommandAction
}

func NewUserService(userCommandAction cmdaction.UserCommandAction) UserService {
	return UserService{
		UserCommandAction: userCommandAction,
	}
}

func (s *UserService) Save() {

}
