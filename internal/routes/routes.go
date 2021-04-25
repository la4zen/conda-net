package routes

import (
	"github.com/la4zen/conda-net/internal/store"
	"github.com/la4zen/conda-net/internal/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	a := e.Group("/api/users")
	a.Use(middleware.JWT(util.Key))
	a.POST("/accessible", routes.Accessible)
	a.POST("/get", routes.GetUser)
	a.POST("/api/friend/add", routes.FriendRequest)
	return nil
}
