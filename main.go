package main

import (
	"github.com/anraku/chat/infrastructure"
	"github.com/anraku/chat/repository"
	"github.com/anraku/chat/usecase"
	"github.com/jinzhu/gorm"
)

// DB is database connection
var DB *gorm.DB

func main() {
	// Setup db
	db, err := infrastructure.Connect()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	userRepo := repository.NewUserRepository(DB)
	roomRepo := repository.NewRoomRepository(DB)
	messageRepo := repository.NewMessageRepository(DB)

	userInteractor := usecase.NewUserInteractor(userRepo, messageRepo)
	roomInteractor := usecase.NewRoomInteractor(roomRepo, messageRepo)
	messageInteractor := usecase.NewMessageInteractor(messageRepo)

	app := NewRouter(userInteractor, roomInteractor, messageInteractor)
	app.Start(":8080")
}
