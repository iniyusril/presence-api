package cmdservice

import (
	"github.com/iniyusril/presence-api/api/server/srvresponse"
	"github.com/iniyusril/presence-api/command/cmdaction"
	"github.com/iniyusril/presence-api/entity"
	"github.com/iniyusril/presence-api/util"
	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
)

type UserService struct {
	UserCommandAction cmdaction.UserCommandAction
	Db                *gorm.DB
}

func NewUserService(userCommandAction cmdaction.UserCommandAction, db *gorm.DB) UserService {
	return UserService{
		UserCommandAction: userCommandAction,
		Db:                db,
	}
}

func (s *UserService) Save(user entity.User) *srvresponse.UserResponse {
	s.Db.Begin()
	defer util.CommitOrRollback(s.Db)
	res, err := s.UserCommandAction.Save(user)
	util.PanicIfError(err)

	srvUserResponse := srvresponse.UserResponse{}
	automapper.Map(res, &srvUserResponse)
	return &srvUserResponse
}

