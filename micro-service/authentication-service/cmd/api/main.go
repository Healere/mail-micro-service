package main

import (
	"authentication/cmd/api/common"
	"authentication/cmd/api/route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = route.CollectRouter(r)
	port := viper.GetString("server.port")
	log.Println("port is: " + port)
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()

	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	//viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
