package route

import (
	"github.com/iniyusril/presence-api/api/server"
	"github.com/iniyusril/presence-api/exception"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute(userServer server.UserServer) *echo.Echo {
	e := echo.New()

	userRoute := e.Group("/api/user")
	userRoute.POST("/", userServer.Save)
	userRoute.GET("/", userServer.FindAll)

	e.HTTPErrorHandler = exception.ErrorHandler

	e.Use(
		middleware.Recover(),   // Recover from all panics to always have your server up
		middleware.Logger(),    // Log everything to stdout
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)

	return e
}
