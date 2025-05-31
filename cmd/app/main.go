package main

import (
	"citatnik/config"
	"citatnik/internal/app"
)

func main() {
	cfg := config.MustLoad()
	app.Run(cfg)
}
