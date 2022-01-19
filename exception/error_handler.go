package exception

import (
	"net/http"

	"github.com/iniyusril/presence-api/api/server/srvresponse"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {

	// if notFoundError(err, c) {
	// 	return
	// }

	// if validationErrors(err, c) {
	// 	return
	// }

	internalServerError(err, c)
}

// func validationErrors(err error, c echo.Context) bool {
// 	exception, ok := err.(validator.ValidationErrors)

// 	if ok {

// 		c.Response().Header().Set("Content-Type", "application/json")

// 		webResponse := web.WebResponse{
// 			Code:   http.StatusNotFound,
// 			Status: "BAD REQUEST",
// 			Data:   exception.Error(),
// 		}
// 		c.JSON(http.StatusBadRequest, webResponse)

// 		return true
// 	} else {
// 		return false
// 	}

// }

// func notFoundError(err error, c echo.Context) bool {
// 	ok := strings.Contains(err.Error(), "not found")

// 	if ok {
// 		c.Response().Header().Set("Content-Type", "application/json")

// 		webResponse := web.WebResponse{
// 			Code:   http.StatusNotFound,
// 			Status: "NOT FOUND",
// 			Data:   err.Error(),
// 		}
// 		c.JSON(http.StatusNotFound, webResponse)

// 		return true
// 	} else {
// 		return false
// 	}

// }

func internalServerError(err error, c echo.Context) bool {

	c.Response().Header().Set("Content-Type", "application/json")

	webResponse := srvresponse.Response{
		Success:   false,
		Message:   err.Error(),
		ErrorCode: 500,
		Data:      nil,
	}
	c.JSON(http.StatusInternalServerError, webResponse)
	return true
}
