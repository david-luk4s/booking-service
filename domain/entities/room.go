package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name          string
	ID            int32
	Level         int32
	InMaintenance bool
	IsAvailable   bool
}
