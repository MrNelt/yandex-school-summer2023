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
	Cost          int32    `json:"cost"`
	CompletedTime string   `json:"completed_time"`
}

type CompleteInfo struct {
	CourierId    int64  `json:"courier_id"`
	OrderId      int64  `json:"order_id" gorm:"primaryKey"`
	CompleteTime string `json:"complete_time"`
}

type OrderDB struct {
	ID         int64 `gorm:"primaryKey" `
	Attributes datatypes.JSON
}

type CourierDB struct {
	ID         int64 `gorm:"primaryKey"`
	Attributes datatypes.JSON
}
