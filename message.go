package main

import "time"

// Mssageは1つのメッセージを表します。
type Message struct {
	ID       int       `gorm:"AUTO_INCREMENT;column:id"`
	UserName string    `gorm:"type:varchar(128);column:user_name"`
	RoomID   int       `gorm:"column:room_id"`
	Message  string    `gorm:"type:varchar(255);column:message"`
	When     time.Time `gorm:"-"`
}
