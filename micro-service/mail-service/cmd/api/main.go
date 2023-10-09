package main

import (
	"log"
	"mailer-service/cmd/api/model"
	"mailer-service/cmd/api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Mailer model.Mail
}

func main() {
	InitConfig()
	webPort := viper.GetString("ports.webPort")
	// connect to database
	r := gin.Default()
	r = routes.Collection(r)
	log.Println("Starting mail service on port", webPort)
	if webPort != "" {
		panic(r.Run(":" + webPort))
	}

	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
