package routes

import (
	"github.com/la4zen/conda-net/internal/model"
	"github.com/labstack/echo"
)

func (r *Routes) Login(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)
	if user.Login == "" || user.Password == "" {
		return c.String(400, "login and password required")
	}
	user, err := r.store.User.GetUser(user)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, user)
}

func (r *Routes) Register(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)
	user, err := r.store.User.CreateUser(user)
	if err != nil {
		return c.String(400, err.Error())
	}
	return c.JSON(200, user)
}
