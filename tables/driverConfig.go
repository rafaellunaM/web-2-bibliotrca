package tables

import (
	"gorm.io/gorm"
)

type Librarian struct {
	gorm.Model
	CPF         string  `gorm:"primaryKey;column:cpf"`
	Name        string  `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	Password    string  `gorm:"column:password"`
	PhoneNumber *string `gorm:"column:phone_number"`
}

func (Librarian) TableName() string {
	return "librarian"
}
