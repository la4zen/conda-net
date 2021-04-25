package routes

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/la4zen/conda-net/internal/model"
	"github.com/labstack/echo"
)

func (r *Routes) FriendRequest(c echo.Context) error {
	friend := &model.Friend{}
	user := &model.User{
		LastLogin: time.Now(),
	}
	c.Bind(&friend)
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(claims["id"].(string))
	user.ID = uint(id)
	if friend.FromUser == user.ID {
		return c.String(400, "you cannot add yourself")
	}
	friend.FromUser = user.ID
	response := r.store.User.AddFriend(friend)
	if response.Err != nil {
		return c.String(response.Code, response.Err.Error())
	}
	return c.NoContent(response.Code)
}
