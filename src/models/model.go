package models

import (
	"gorm.io/datatypes"
)

type Courier struct {
	CourierID    int64    `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int    `json:"regions"`
	WorkingHours []string `json:"working_hours"`
}

type CourierDB struct {
	ID         int64 `gorm:"primaryKey"`
	Attributes datatypes.JSON
}
