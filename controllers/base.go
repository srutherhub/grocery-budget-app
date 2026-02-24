package controllers

import (
	h "app/handlers"

	"github.com/srutherhub/web-app/controller"
)

func BaseController() controller.Controller {
	baseController := controller.New()

	baseController.SetBase("/")

	baseController.RegisterRoute(controller.Route{Method: "GET", Path: "", Handler: h.Base})
	baseController.RegisterRoute(controller.Route{Method: "GET", Path: "app", Handler: h.App})

	return *baseController
}
