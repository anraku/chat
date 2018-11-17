package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type User struct {
// 	ID   int    `gorm:"column:'id'"`
// 	Name string `gorm:"column:'name'"`
// }

func Connect() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	protocol := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	db, err := gorm.Open("mysql", user+":"+pass+"@"+protocol+"/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	//db.AutoMigrate(&User{})
	return db, err
}
