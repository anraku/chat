package main

import "time"

// Mssageは1つのメッセージを表します。
type Message struct {
	ID       int       `gorm:"column:id"`
	UserID   int       `gorm:"column:user_id"`
	RoomID   int       `gorm:"column:room_id"`
	Message  string    `gorm:"column:message"`
	UserName string    `gorm:"-"`
	When     time.Time `gorm:"-"`
}
