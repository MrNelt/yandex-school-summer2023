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

type Order struct {
	OrderID       int64    `json:"order_id"`
	Weight        int      `json:"weight"`
	Regions       int      `json:"regions"`
	DeliveryHours []string `json:"delivery_hours"`
	Cost          int      `json:"cost"`
}

type CourierDB struct {
	ID         int64 `gorm:"primaryKey"`
	Attributes datatypes.JSON
}
