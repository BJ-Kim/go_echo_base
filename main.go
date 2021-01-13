package main

import (
	"os"
	"web_base/app"
	"web_base/migrate"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		mig := migrate.NewMigrate()
		switch argsWithoutProg[0] {
		case "upgrade":
			mig.Upgrade()
		case "downgrade":
			mig.Downgrade()
		}
	} else {
		server := app.NewServer()
		server.RunServer()
	}
}
