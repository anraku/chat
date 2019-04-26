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
	roomRepo := datastore.NewRoomMySQLRepository(db)
	messageRepo := datastore.NewMessageMySQLRepository(db)

	messageService := service.NewMessageService(messageRepo)

	userInteractor := usecase.NewUserUsecase(userRepo)
	roomInteractor := usecase.NewRoomUsecase(roomRepo)
	messageInteractor := usecase.NewMessageUsecase(messageService, messageRepo)

	app := router.NewRouter(userInteractor, roomInteractor, messageInteractor)
	app.Start(":8080")
}
