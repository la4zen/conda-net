package routes

import (
	"time"

	"github.com/la4zen/conda-net/internal/model"
	"github.com/la4zen/conda-net/internal/util"
	"github.com/labstack/echo"
)

func (r *Routes) FriendRequest(c echo.Context) error {
	friend := &model.Friend{}
	user := &model.User{
		LastLogin: time.Now(),
	}
	c.Bind(&friend)
	if util.VerifyToken(user, c.Request().Header.Get("Authorization")) != nil {
		return c.NoContent(401)
	}
	if friend.FromUser == user.ID {
		return c.String(400, "you dont add yourself")
	}
	friend.FromUser = user.ID
	response := r.store.User.AddFriend(friend)
	if response.Err != nil {
		return c.String(response.Code, response.Err.Error())
	}
	return c.NoContent(response.Code)
}
