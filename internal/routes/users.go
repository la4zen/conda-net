package routes

import (
	"github.com/la4zen/conda-net/internal/model"
	"github.com/labstack/echo"
)

func (r *Routes) GetUser(c echo.Context) error {
	user := &model.User{}
	c.Bind(user)
	if user.ID == 0 || user.Login == "" {
		return c.String(400, "id or login required")
	}
	response := r.store.User.GetUser(user)
	if response.Err != nil {
		return c.String(response.Code, response.Err.Error())
	}
	return c.JSON(response.Code, map[string]interface{}{
		"user": user,
	})
}
