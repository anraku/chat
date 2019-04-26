package main

import (
	"github.com/anraku/chat/domain/service"
	"github.com/anraku/chat/infrastructure/datastore"
	"github.com/anraku/chat/infrastructure/router"
	"github.com/anraku/chat/usecase"
)

func main() {
	// Setup db
	db, err := datastore.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := datastore.NewUserSessionRepository()
	roomRepo := datastore.NewRoomRepository(db)
	messageRepo := datastore.NewMessageRepository(db)

	messageService := service.NewMessageService(messageRepo)

	userInteractor := usecase.NewUserInteractor(userRepo)
	roomInteractor := usecase.NewRoomInteractor(roomRepo, messageRepo)
	messageInteractor := usecase.NewMessageInteractor(messageService, messageRepo)

	app := router.NewRouter(userInteractor, roomInteractor, messageInteractor)
	app.Start(":8080")
}
