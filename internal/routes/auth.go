package routes

import (
	"time"

	"github.com/la4zen/conda-net/internal/model"
	"github.com/la4zen/conda-net/internal/util"
	"github.com/labstack/echo"
)

func (r *Routes) Login(c echo.Context) error {
	user := &model.User{
		LastLogin: time.Now(),
	}
	c.Bind(&user)
	if !util.CheckFields(user.Login) {
		return c.String(400, "login and password required")
	}
	util.MD5Password(&user.Password)
	response := r.store.User.FindUser(user)
	if response.Err != nil {
		return c.String(response.Code, response.Err.Error())
	}
	token, err := util.CreateToken(user)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(response.Code, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (r *Routes) Register(c echo.Context) error {
	user := &model.User{
		LastLogin: time.Now(),
	}
	c.Bind(&user)
	if !util.CheckFields(user.FirstName, user.LastName, user.Login, user.Password) {
		return c.String(400, "all fields required")
	}
	util.MD5Password(&user.Password)
	response := r.store.User.CreateUser(user)
	if response.Err != nil {
		return c.String(response.Code, response.Err.Error())
	}
	token, err := util.CreateToken(user)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(response.Code, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
