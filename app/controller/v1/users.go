package v1

import (
	"github.com/labstack/echo/v4"
)

type usersController struct {
	ctx *echo.Context
}

func NewUsersController(ctx *echo.Context) *usersController {
	return &usersController{ctx}
}

func (uc *usersController) InsertUser() {
	println("INSERT USER")
}
