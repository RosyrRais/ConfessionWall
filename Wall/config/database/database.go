package database

import (
	"fmt"

	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	//user := config.Config.GetString("database.user")
	user := "root"
	//	pass := config.Config.GetString("database.pass")
	pass := "123456"
	//	port := config.Config.GetString("database.port")
	port := 3306
	//	host := config.Config.GetString("database.host")
	host := "127.0.0.1"
	DBname := "wall"
	//	log.Fatal(pass)
	//dsn := fmt.Sprintf("host=%s user=%s pass=%s DBname=%s post=%s", host, user, pass, DBname, port)
	//dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&&parseTime=True&loc=Local", user, pass, host, port, DBname)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, DBname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	err = autoMigrate(db)
	if err != nil {
		log.Fatal("表单创建失败:", err)
	}
	DB = db
}
