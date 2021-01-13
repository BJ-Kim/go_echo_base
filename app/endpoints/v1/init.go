package v1

import (
	"github.com/labstack/echo/v4"
	"web_base/app/controller"
)

type v1Endpoints struct {
	usersEndpoints *usersEndpoints
}

func NewV1Endpoints(api *echo.Group) *v1Endpoints {
	controller := controller.NewController()
	return &v1Endpoints{
		usersEndpoints: newUsersEndpoints(controller, api.Group("/v1")),
	}
}

func (e *v1Endpoints) Register() {
}
