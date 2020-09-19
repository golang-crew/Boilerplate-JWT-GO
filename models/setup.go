package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var gGormDB *gorm.DB

func CloseDB() {
	if gGormDB != nil {
		gGormDB.Close()
	}
}

func Init() {
	db, err := gorm.Open("mysql", dbConnString())
	if err != nil {
		log.Println("[DB Error] ", err)
		panic(err)
	}
	gGormDB = db

	log.Println("[DB] Start DB Migration ... ")
	log.Println("[DB] Start DB ... ")

	if err := db.AutoMigrate(&Users{}, &Todo{}).Error; err != nil {
		log.Println("[DB Error] ", err)
		panic(err)
	}
}

func dbConnString() (dbConnString string) {
	dbHost := viper.GetString(`database.host`)
	dbHost = "127.0.0.1"
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	dbConnString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName,
	)

	return
}
