package models

import (
	"gorm.io/gorm"
)

type Users struct {
	DefaultModel
	// Ts        *transaction
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
}

func (u *Users) InsertUser(tx *gorm.DB) error {
	return tx.Create(u).Error
}

// func (u *Users) Testt() {
//     u.Ts.Tx.Create(u)
// }
