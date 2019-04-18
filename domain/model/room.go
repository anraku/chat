package model

import (
	"github.com/anraku/chat/infrastructure/trace"
)

// Roomは一つのチャットルームを表します
type Room struct {
	ID          int    `gorm:"AUTO_INCREMENT;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	// Forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	Forward chan *Message `gorm:"-"`
	// Joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	Join chan *User `gorm:"-"`
	// Leaveはチャットルームから退室しようとしているクライアントのためのチャネルです
	Leave chan *User `gorm:"-"`
	// Usersには在室しているすべてのクライアントが保持されます。
	Users map[*User]bool `gorm:"-"`
	// Tracerはチャットルーム上で行われた操作のログを受け取ります。
	Tracer trace.Tracer `gorm:"-"`
}
