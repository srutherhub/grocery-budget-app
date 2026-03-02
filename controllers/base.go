package controllers

import (
	h "app/handlers"
	m "app/middleware"
	s "app/services"

	"github.com/srutherhub/web-app/controller"
)

func BaseController(services s.AppServices) controller.Controller {
	baseController := controller.New()

	baseController.SetBase("/")

	baseController.RegisterRoute(controller.Route{Method: "GET", Path: "", Handler: h.Base()})
	baseController.RegisterRoute(controller.Route{Method: "GET", Path: "app", Handler: m.IsLoggedIn(services.Auth)(h.App())})
	baseController.RegisterRoute(controller.Route{Method: "GET", Path: "login", Handler: h.Login()})
	baseController.RegisterRoute(controller.Route{Method: "GET", Path: "signup", Handler: h.Signup()})

	return *baseController
}
