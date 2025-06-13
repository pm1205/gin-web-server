package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username" binding:"required,min=3,max=32"`
	Password  string    `json:"password" binding:"required,min=6,max=32"`
	Avatar    string    `json:"avatar"`
	Status    int       `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
