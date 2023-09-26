package database

import (
	"Wall/config/config"

	"fmt"

	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	user := config.Config.Getstring("database.user")
	pass := config.Config.Getstring("database.pass")
	port := config.Config.Getstring("database.port")
	host := config.Config.Getstring("database.host")
	DBname := config.Config.Getstring("database.DBname")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf&&parseTime=True&loc=Local", user, pass, host, port, DBname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("database connect error", err)
	}
	err = autoMigrate(db)
	if err != nil {
		log.Fatal("table create error", err)
	}
	DB = db
}
