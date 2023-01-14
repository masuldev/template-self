package main

import (
	"github.com/masuldev/template-self/config"
	"github.com/masuldev/template-self/internal/app"
)

func main() {
	env := config.GetAll()

	app.Run(env)
}
