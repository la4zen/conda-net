package main

import "github.com/la4zen/conda-net/internal/app"

const configpath string = "configs/config.json"

func main() {
	app.Run(configpath)
}
