package routes

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/la4zen/conda-net/internal/model"
	"github.com/labstack/echo"
)

func (r *Routes) GetUser(c echo.Context) error {
	user := &model.User{}
	c.Bind(user)
	if user.ID == 0 || user.Login == "" {
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		user.ID = uint(claims["ID"].(float64))
	}
	response := r.store.User.GetUser(user)
	if response.Err != nil {
		return c.String(response.Code, response.Err.Error())
	}
	return c.JSON(response.Code, map[string]interface{}{
		"user": user,
	})
}
