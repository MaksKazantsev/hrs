package main

import (
	"github.com/alserov/hrs/auth/internal/app"
	"github.com/alserov/hrs/auth/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
