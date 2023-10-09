package RPC

import (
	"context"
	"log-service/cmd/api/logs"
	"log-service/cmd/api/model"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models model.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := model.LogEntry{
		Email: input.Email,
		Data:  input.Data,
	}

	err := model.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

// func gRPCListen() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
// 	if err != nil {
// 		log.Fatalf("Failed to listen for gRPC: %v", err)
// 	}

// 	s := grpc.NewServer()

// 	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

// 	log.Printf("gRPC Server started on port %s", gRpcPort)

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("Failed to listen for gRPC: %v", err)
// 	}
// }
