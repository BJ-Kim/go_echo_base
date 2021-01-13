package endpoints

import (
	"github.com/labstack/echo/v4"
	"web_base/app/endpoints/v1"
)

type endpoints struct {
	api *echo.Group
}

func NewEndpoints(e *echo.Echo) *endpoints {
	return &endpoints{e.Group("/api")}
}

func (ep *endpoints) RegisterEndpoints() {
	ep.registerV1Endpoints()
}

func (ep *endpoints) registerV1Endpoints() {
	v1.NewV1Endpoints(ep.api)
}
