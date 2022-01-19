//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/iniyusril/presence-api/api/server"
	"github.com/iniyusril/presence-api/command/cmdaction"
	"github.com/iniyusril/presence-api/command/cmdservice"
	"github.com/iniyusril/presence-api/config"
	"github.com/iniyusril/presence-api/query/qryaction"
	"github.com/iniyusril/presence-api/query/qryservice"
	"github.com/iniyusril/presence-api/repository"
	"github.com/iniyusril/presence-api/route"
	"github.com/labstack/echo/v4"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	cmdaction.NewUserCommandAction,
	cmdservice.NewUserService,
	qryaction.NewUserQueryAction,
	qryservice.NewUserQueryService,
	server.NewUserServer,
)

func InitializedServer() *echo.Echo {
	wire.Build(config.NewDb, route.NewRoute, userSet)
	return nil
}
