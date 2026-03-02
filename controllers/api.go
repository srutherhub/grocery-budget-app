package controllers

import (
	h "app/handlers"
	s "app/services"

	"github.com/srutherhub/web-app/controller"
)

func ApiController(services s.AppServices) controller.Controller {
	apiController := controller.New()

	apiController.SetBase("/api")

	apiController.RegisterRoute(controller.Route{Method: "POST", Path: "/chat/submit", Handler: h.ApiSubmitChat()})

	apiController.RegisterRoute(controller.Route{Method: "POST", Path: "/auth/login", Handler: h.ApiLogin(services.Auth)})
	apiController.RegisterRoute(controller.Route{Method: "POST", Path: "/auth/signup", Handler: h.ApiSignup(services.Auth)})

	return *apiController
}
