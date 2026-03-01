package controllers

import (
	h "app/handlers"

	"github.com/srutherhub/web-app/controller"
)

func ApiController() controller.Controller {
	apiController := controller.New()

	apiController.SetBase("/api")

	apiController.RegisterRoute(controller.Route{Method: "POST", Path: "/chat/submit", Handler: h.SubmitChat()})

	return *apiController
}
