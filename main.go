package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-crew/Bolierplate-JWT-Auth/common"
	"github.com/golang-crew/Bolierplate-JWT-Auth/models"
	"github.com/golang-crew/Bolierplate-JWT-Auth/routers"
	"github.com/spf13/viper"
)

func viperInit() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	viperInit()
	common.RedisInit()
	models.Init()
	defer models.CloseDB()
	routers.Init(r)

	log.Fatal(r.Run(":8080"))
}
