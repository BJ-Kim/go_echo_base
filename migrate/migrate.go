package migrate

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"web_base/constants"
)

// how to create migrate sql file:
// "migrate create -ext sql -dir migrate/migrations -seq {foo}"
type dbMigrate struct {
	m *migrate.Migrate
}

func NewMigrate() *dbMigrate {
	dc := constants.NewDbConstants()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=public", dc.User, dc.Password, dc.Host, dc.Port, dc.Database)
	m, err := migrate.New(
		"file://migrate/migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &dbMigrate{m}
}

func (db *dbMigrate) Upgrade() {
	defer func() {
		db.m.Close()
	}()
	if err := db.m.Up(); err != nil {
		log.Fatal(err)
	}
}

func (db *dbMigrate) Downgrade() {
	defer func() {
		db.m.Close()
	}()
	if err := db.m.Down(); err != nil {
		log.Fatal(err)
	}
}
