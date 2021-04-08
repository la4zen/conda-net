package app

import (
	"github.com/la4zen/conda-net/internal/store"
	"github.com/labstack/echo"
)

func Run() {
	store, err := store.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Logger.Fatal(e.Start(store.Config.Http.IP + ":" + store.Config.Http.Port))
}
