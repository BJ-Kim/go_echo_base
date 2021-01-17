package v1

import (
	"web_base/app/models"
	"web_base/app/modules"
)

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (uc *UsersController) InsertUser(users *models.Users) error {
	ts := models.NewTransaction()
	defer func() {
		if r := recover(); r != nil {
			ts.RollbackTransaction()
		}
	}()
	if err := users.InsertUser(ts.GetTransaction()); err != nil {
		ts.RollbackTransaction()
		return modules.NewInvalidValueError()
	}
	ts.CommitTransaction()
	return nil
}
