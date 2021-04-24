package routes

import (
	"github.com/la4zen/conda-net/internal/store"
	"github.com/labstack/echo"
)

type Routes struct {
	store *store.Store
}

func New(store *store.Store) *Routes {
	return &Routes{
		store: store,
	}
}

func Set(e *echo.Echo, routes *Routes) error {
	e.POST("/api/login", routes.Login)
	e.POST("/api/register", routes.Register)
	e.POST("/api/user", routes.GetUser)
	e.POST("/api/friend/add", routes.FriendRequest)
	return nil
}
