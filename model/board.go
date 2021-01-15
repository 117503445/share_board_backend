package model

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Data string
}
