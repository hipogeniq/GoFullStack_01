package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `gorm:"text; not null"`
	Answer   string `gorm:"text;not null;"`
}
