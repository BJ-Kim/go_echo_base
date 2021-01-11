package v1

import (
	"github.com/labstack/echo/v4"
	"web_base/app/controller"
	"web_base/app/models"
)

type usersController struct {
	controller.Controller
	ctx *echo.Context
}

func NewUsersController(ctx *echo.Context) *usersController {
	return &usersController{ctx: ctx}
}

func (uc *usersController) InsertUser(users *models.Users) {
	println("INSERT USER")
}
