package models

import "gorm.io/gorm"

type Courier struct {
	ID string `json:"id" query:"id"`
}

type Test struct {
	gorm.Model
}
