package qryservice

import (
	"github.com/iniyusril/presence-api/api/server/srvresponse"
	"github.com/iniyusril/presence-api/query/qryaction"
	"github.com/iniyusril/presence-api/util"
	"github.com/stroiman/go-automapper"
)

type UserQueryService struct {
	UserQueryAction qryaction.UserQueryAction
}

func NewUserQueryService(userQueryAction qryaction.UserQueryAction) UserQueryService {
	return UserQueryService{
		UserQueryAction: userQueryAction,
	}
}

func (s *UserQueryService) FindAll() *[]srvresponse.UserResponse {
	res, err := s.UserQueryAction.FindAll()
	util.PanicIfError(err)
	response := make([]srvresponse.UserResponse, 0)
	automapper.Map(*res, &response)
	return &response
}
