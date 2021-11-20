package model

import "time"

type CreditCard struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `gorm:"user_id" json:"user_id"`
	Number    string     `json:"number"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (CreditCard) TableName() string { return "credit_cards" }
