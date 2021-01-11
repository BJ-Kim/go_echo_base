package v1

import (
	"github.com/labstack/echo/v4"
)

type v1Endpoints struct {
	v1 *echo.Group
}

func NewV1Endpoints(api *echo.Group) *v1Endpoints {
	return &v1Endpoints{api.Group("/v1")}
}

func (e *v1Endpoints) Register() {
	newUsersEndpoints(e.v1)
}
