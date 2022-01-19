package server

import (
	"net/http"

	"github.com/iniyusril/presence-api/api/server/srvrequest"
	"github.com/iniyusril/presence-api/api/server/srvresponse"
	"github.com/iniyusril/presence-api/command/cmdservice"
	"github.com/iniyusril/presence-api/entity"
	"github.com/iniyusril/presence-api/query/qryservice"
	"github.com/iniyusril/presence-api/util"

	"github.com/labstack/echo/v4"
)

type UserServer struct {
	UserCommandService cmdservice.UserService
	UserQueryService   qryservice.UserQueryService
}

func NewUserServer(UserCommandService cmdservice.UserService, userQuerService qryservice.UserQueryService) UserServer {

	return UserServer{
		UserCommandService: UserCommandService,
		UserQueryService:   userQuerService,
	}
}

func (s *UserServer) Save(c echo.Context) error {
	userRequest := new(srvrequest.UserRequest)
	if err := c.Bind(userRequest); err != nil {
		util.PanicIfError(err)
	}
	userEntity := entity.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}

	res := s.UserCommandService.Save(userEntity)

	response := srvresponse.Response{
		Success: true,
		Message: "user has been saved",
		Data:    res,
	}
	return c.JSON(http.StatusOK, response)
}

func (s *UserServer) FindAll(c echo.Context) error {

	res := s.UserQueryService.FindAll()

	response := srvresponse.Response{
		Success: true,
		Message: "user has been saved",
		Data:    res,
	}
	return c.JSON(http.StatusOK, response)
}
