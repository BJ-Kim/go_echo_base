package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"web_base/constants"
)

type DefaultModel struct {
	ID        uint      `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	RemovedAt time.Time `gorm:"column:removed_at" json:"removed_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type dbServer struct {
	db *gorm.DB
}

var dbServerInstance *dbServer

func GetDbServerInstance() *dbServer {
	if dbServerInstance == nil {
		dbServerInstance = new(dbServer)
		dc := constants.NewDbConstants()
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", dc.Host, dc.User, dc.Password, dc.Database, dc.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		dbServerInstance.db = db
	}
	return dbServerInstance
}

type transaction struct {
	tx *gorm.DB
}

func NewTransaction() *transaction {
	dbServer := GetDbServerInstance()
	tx := dbServer.db.Begin()
	return &transaction{tx: tx}
}

func (t *transaction) GetTransaction() *gorm.DB {
	return t.tx
}

func (t *transaction) RollbackTransaction() {
	t.tx.Rollback()
}

func (t *transaction) CommitTransaction() {
	t.tx.Commit()
}
