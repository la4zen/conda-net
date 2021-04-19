package routes

import (
	"time"

	"github.com/la4zen/conda-net/internal/model"
	"github.com/la4zen/conda-net/internal/util"
	"github.com/labstack/echo"
)

func (r *Routes) Login(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)
	if user.Login == "" || user.Password == "" {
		return c.String(400, "login and password required")
	}
	user, code, err := r.store.User.GetUser(user)
	if err != nil {
		return c.String(code, err.Error())
	}
	return c.JSON(code, user)
}

func (r *Routes) Register(c echo.Context) error {
	user := &model.User{
		LastLogin: time.Now(),
	}
	c.Bind(&user)
	if !util.AllTrue(user.FirstName, user.LastName, user.Login, user.Password) {
		return c.String(400, "all fields required")
	}
	util.MD5Password(&user.Password)
	user, code, err := r.store.User.CreateUser(user)
	if err != nil {
		return c.String(code, err.Error())
	}
	return c.JSON(code, user)
}
