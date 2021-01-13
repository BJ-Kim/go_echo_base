package v1

import (
	"web_base/app/models"
)

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (uc *UsersController) InsertUser(users *models.Users) {
	println("INSERT USER")
}
