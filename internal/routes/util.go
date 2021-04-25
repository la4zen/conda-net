package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func (r *Routes) Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
