package main

import (
	c "app/controllers"

	"github.com/srutherhub/web-app/server"
)

func main() {
	cfg := server.ServerConfig{Port: "5555"}

	s := server.New()

	s.RegisterController(c.BaseController())
	s.RegisterController(c.ApiController())

	s.Start(cfg)
}
