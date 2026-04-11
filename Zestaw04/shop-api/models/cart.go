package models

type Cart struct {
	ID    uint       `json:"id" gorm:"primaryKey"`
	Items []CartItem `json:"items"`
}