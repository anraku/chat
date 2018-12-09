package main

import (
	"github.com/anraku/chat/database"
	"github.com/jinzhu/gorm"
)

// DB is database connection
var DB *gorm.DB

func main() {
	// Setup db
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	userRepo := NewUserRepository()
	roomRepo := NewRoomRepository(DB)
	messageRepo := NewMessageRepository(DB)

	userInteractor := NewUserInteractor(userRepo, messageRepo)
	roomInteractor := NewRoomInteractor(roomRepo, messageRepo)
	messageInteractor := NewMessageInteractor(messageRepo)

	app := NewRouter(userInteractor, roomInteractor, messageInteractor)
	app.Start(":8080")
}
