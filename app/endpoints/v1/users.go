package v1

import (
	"github.com/labstack/echo/v4"
	v1c "web_base/app/controller/v1"
)

type usersEndpoints struct {
	users *echo.Group
}

func newUsersEndpoints(v1 *echo.Group) {
	ue := usersEndpoints{v1.Group("/users")}
	{
		ue.users.GET("", ue.getUser)
		ue.users.POST("", ue.insertUser)
	}
}

func (ue *usersEndpoints) getUser(ctx echo.Context) error {
	return ctx.String(200, "Hello, World!")
}

func (ue *usersEndpoints) insertUser(ctx echo.Context) error {
	uc := v1c.NewUsersController(&ctx)
	uc.InsertUser()
	return ctx.String(200, "Hello, World!")
}
