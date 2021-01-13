package constants

import (
	"os/user"
)

type dbConstants struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDbConstants() *dbConstants {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return &dbConstants{
		Host:     "localhost",
		Port:     "5432",
		User:     user.Username,
		Password: "asdqwe123",
		Database: "forganization",
	}
}

func (dc *dbConstants) GetDatabaseInfo() {
	//TODO: aws secret
}
