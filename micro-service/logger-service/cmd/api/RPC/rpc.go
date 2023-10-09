package RPC

import (
	"context"
	"log"
	"log-service/cmd/api/common"
	"log-service/cmd/api/model"
	"time"
)

// RPCServer is the type for our RPC Server. Methods that take this as a receiver are available
// over RPC, as long as they are exported.
type RPCServer struct{}

// RPCPayload is the type for data we receive from RPC
type RPCPayload struct {
	Email string
	Data  string
}

// LogInfo writes our payload to mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	client, err := common.ConnectToMongo()
	if err != nil {
		return err
	}

	collection := client.Database("logs").Collection("logs")
	_, err = collection.InsertOne(context.TODO(), model.LogEntry{
		Email:     payload.Email,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error writing to mongo: ", err)
		return err
	}

	// resp is the message sent back to the RPC caller
	*resp = "Processed payload via RPC:" + payload.Email
	return nil
}
