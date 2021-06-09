package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	ID      int
	Title   string
	Address string
}
