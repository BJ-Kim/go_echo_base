package v1

import ()

type V1Controller struct {
	UsersController *UsersController
}

func NewV1Controller() *V1Controller {
	return &V1Controller{
		UsersController: NewUsersController(),
	}
}
