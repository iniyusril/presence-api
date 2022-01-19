package route

import (
	"github.com/iniyusril/presence-api/api/server"
	"github.com/labstack/echo/v4"
)

func NewRoute(userServer server.UserServer) *echo.Echo {
	e := echo.New()

	userRoute := e.Group("/api/user")
	userRoute.POST("/", userServer.Save)
	userRoute.GET("/", userServer.FindAll)

	return e
}
