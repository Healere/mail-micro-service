package controller

import (
	"fmt"
	"log"
	"log-service/cmd/api/model"
	"log-service/cmd/api/util"
	"net"
	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type JSONPayload struct {
	Email   string `json:"email"`
	LogData string `json:"logData"`
}

func WriteLog(c *gin.Context) {
	// read json into var
	var requestPayload JSONPayload
	c.ShouldBind(&requestPayload)

	fmt.Println("request is: ", requestPayload)
	// insert data
	event := model.LogEntry{
		Email: requestPayload.Email,
		Data:  requestPayload.LogData,
	}
	err := model.Insert(event)
	if err != nil {
		util.ErrorJSON(c.Writer, err)
		return
	}

	resp := util.JsonResponse{
		Error:   false,
		Message: "logged",
	}

	util.WriteJSON(c.Writer, http.StatusAccepted, resp)
}

func RpcListen() error {
	rpcPort := viper.GetString("port.rpcPort")
	log.Println("Starting RPC server on port: ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	rpcConn, err := listen.Accept()
	for {
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
