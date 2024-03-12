package structur

import "time"

type Orders struct {
	OrderID      int `json:"PrimaryKey" gorm:"primaryKey"`
	OrderedAt    time.Time
	CustomerName string
	Items        []Items `gorm:"foreignKey:OrderID"`
}

type Items struct {
	ItemID      int    `json:"PrimaryKey" gorm:"primaryKey"`
	ItemCode    string `json:"ItemCode"`
	Description string `json:"Description"`
	Quantity    int    `json:"Quantity"`
	OrderID     int    `gorm:"index"` // Foreign key
}
