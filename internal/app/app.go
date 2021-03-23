package app

import (
	"github.com/la4zen/conda-net/internal/config"
	"github.com/la4zen/conda-net/internal/routes"
	"github.com/labstack/echo/v4"
)

func Run(configpath string) {
	var config config.Config = config.NewConfig(configpath)
	e := echo.New()
	e.POST("/v1/auth/login", routes.Login)
	e.POST("/v1/auth/register", routes.Register)
	e.Logger.Fatal(e.Start(config.Http.IP + ":" + config.Http.Port))
}
