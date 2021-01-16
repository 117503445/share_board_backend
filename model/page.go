package model

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	PageNumber int
	BoardID    string
	Data       string
}
