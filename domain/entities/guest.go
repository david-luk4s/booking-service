package entities

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
}

func NewGuest() *Guest {
	return &Guest{}
}
