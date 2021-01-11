package models

type Users struct {
	DefaultModel
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
}
