package controller

import (
	"web_base/app/controller/v1"
)

type Controller struct {
	V1 *v1.V1Controller
}

func NewController() *Controller {
	controller := Controller{
		V1: v1.NewV1Controller(),
	}
	return &controller
}
