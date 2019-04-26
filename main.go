package main

import (
	"github.com/anraku/chat/domain/repository"
	"github.com/anraku/chat/infrastructure/database/persistence"
	"github.com/anraku/chat/interfaces/api/server/router"
	"github.com/anraku/chat/usecase"
)

func main() {
	// Setup db
	db, err := persistence.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUserSessionRepository()
	roomRepo := repository.NewRoomRepository(db)
	messageRepo := repository.NewMessageRepository(db)

	userInteractor := usecase.NewUserInteractor(userRepo, messageRepo)
	roomInteractor := usecase.NewRoomInteractor(roomRepo, messageRepo)
	messageInteractor := usecase.NewMessageInteractor(messageRepo)

	app := router.NewRouter(userInteractor, roomInteractor, messageInteractor)
	app.Start(":8080")
}
