package model

import (
	"time"
)

type Account struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"column:username;type:VARCHAR(100);NOT NULL"`
	Email     string    `json:"email" gorm:"column:email;type:VARCHAR(100);UNIQUE_INDEX;NOT NULL"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
