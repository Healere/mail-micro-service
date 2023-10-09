package main

import (
	"log"
	"log-service/cmd/api/common"
	"log-service/cmd/api/model"
	"log-service/cmd/api/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Models model.Models
}

func main() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	common.Init()
	// err = rpc.Register(new(RPC.RPCServer))
	// if err != nil {
	// 	log.Println("Error register RPCServer: ", err)
	// }
	// go controller.RpcListen()
	//go app.gRPCListen()
	// start web server
	webPort := viper.GetString("port.webPort")
	log.Println("Starting service on port", webPort)
	r := gin.Default()
	r = route.Collections(r)
	if webPort != "" {
		panic(r.Run(":" + webPort))
	}
	panic(r.Run())

}
