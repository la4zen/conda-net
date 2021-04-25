package app

import (
	"github.com/la4zen/conda-net/internal/routes"
	"github.com/la4zen/conda-net/internal/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	store, err := store.New()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	routes.Set(e, routes.New(store))
	e.Logger.Fatal(e.Start(store.Config.Http.IP + ":" + store.Config.Http.Port))
}
