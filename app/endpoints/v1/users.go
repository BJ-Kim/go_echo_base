package v1

import (
	"github.com/labstack/echo/v4"
	"web_base/app/controller"
	"web_base/app/models"
)

type usersEndpoints struct {
	// v1Endpoints
	controller *controller.Controller
	users      *echo.Group
}

func newUsersEndpoints(controller *controller.Controller, v1 *echo.Group) *usersEndpoints {
	ue := usersEndpoints{
		controller: controller,
		users:      v1.Group("/users"),
	}
	{
		ue.users.GET("", ue.getUser)
		ue.users.POST("", ue.insertUser)
	}
	return &ue
}

func (ue *usersEndpoints) getUser(ctx echo.Context) error {
	return ctx.String(200, "Hello, World!")
}

func (ue *usersEndpoints) insertUser(ctx echo.Context) error {
	users := new(models.Users)
	if err := ctx.Bind(users); err != nil {
	}
	ue.controller.V1.UsersController.InsertUser(users)
	return ctx.String(200, "Hello, World!")
}
