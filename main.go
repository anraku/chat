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

	ur := NewUserRepository()
	ui := NewUserInteractor(ur)
	app := NewRouter(ui)
	app.Start(":8080")
}
