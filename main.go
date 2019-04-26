package main

import (
	"github.com/anraku/chat/infrastructure"
	"github.com/anraku/chat/interfaces/repository"
	"github.com/anraku/chat/usecase"
)

func main() {
	// Setup db
	db, err := infrastructure.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUserSessionRepository()
	roomRepo := repository.NewRoomRepository(db)
	messageRepo := repository.NewMessageRepository(db)

	userInteractor := usecase.NewUserInteractor(userRepo)
	roomInteractor := usecase.NewRoomInteractor(roomRepo, messageRepo)
	messageInteractor := usecase.NewMessageInteractor(messageRepo)

	app := infrastructure.NewRouter(userInteractor, roomInteractor, messageInteractor)
	app.Start(":8080")
}
