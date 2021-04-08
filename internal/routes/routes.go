package routes

import "github.com/labstack/echo"

func Set(e *echo.Echo) {
	e.POST("/api/login", Login)
}
