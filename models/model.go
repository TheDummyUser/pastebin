package models

import "gorm.io/gorm"

type Paste struct {
	gorm.Model
	UUID    string `gorm:"uniqueIndex;not null"` // Make it exported and add constraints
	Content string `gorm:"not null"`
}
