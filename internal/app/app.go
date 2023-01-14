package app

import (
	"fmt"
	"github.com/masuldev/template-self/config"
	"github.com/masuldev/template-self/internal/api"
)

func Run(cfg *config.Env) {
	port := fmt.Sprintf(":%s", cfg.Application.Port)
	route := api.Route()

	err := route.Listen(port)
	if err != nil {
		panic(err)
	}
}

func init() {

}
